package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"

	"github.com/FarrelioGustiana/backend/models"
)

func StartFeedScheduler(db *gorm.DB) {
	c := cron.New()

	_, err := c.AddFunc("@every 15m", func() {
		log.Println("Running scheduled feed fetch job...")
		fetchAndStoreArticles(db)
	})
	
	if err != nil {
		log.Fatalf("Error scheduling feed fetch job: %v", err)
	}

	c.Start()
	log.Println("Feed fetching scheduler started.")
	
	// Run immediately on startup
	log.Println("Running initial feed fetch job...")
	go fetchAndStoreArticles(db)
}

func fetchAndStoreArticles(db *gorm.DB) {

	var feeds []models.Feed
	if err := db.Find(&feeds).Error; err != nil {
		log.Printf("Error fetching feeds for scheduler: %v", err)
		return
	}

	fp := gofeed.NewParser()

	for _, feed := range feeds {

		func(feed models.Feed) {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("Recovered from panic while fetching feed %s: %v", feed.URL, r)
					}
				}()

				ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
				defer cancel()

				rssFeed, err := fp.ParseURLWithContext(feed.URL, ctx)
				if err != nil {
					log.Printf("Error parsing feed %s (%s): %v", feed.Name, feed.URL, err)
					return
				}
				log.Printf("Fetched feed: %s (%s)", rssFeed.Title, feed.URL)

				for _, item := range rssFeed.Items {
					articleLink := item.Link
					guid := item.GUID

					if articleLink == "" && guid == "" {
						log.Printf("Skipping article from %s due to missing link/guid.", feed.URL)
						continue
					}

					if guid == "" {
						guid = articleLink
					}

					var existingArticle models.Article
					if err := db.Where("link = ? OR guid = ?", articleLink, guid).First(&existingArticle).Error; err == nil {
						continue
					} else if !errors.Is(err, gorm.ErrRecordNotFound) {
						log.Printf("Database error checking existing article for feed %s: %v", feed.URL, err)
						continue
					}

					pubDate := time.Now()
					if item.PublishedParsed != nil {
						pubDate = *item.PublishedParsed
					}

					article := models.Article{
						FeedID:      feed.ID,
						Title:       item.Title,
						Link:        articleLink,
						Description: item.Description,
						PubDate:     &pubDate,
						GUID:        guid,
					}

					if err := db.Create(&article).Error; err != nil {
						log.Printf("Error storing article '%s' from feed %s: %v", item.Title, feed.URL, err)
					} else {
						log.Printf("Stored new article: %s", item.Title)
					}
				}
				now := time.Now()
				feed.LastFetchedAt = &now
				db.Save(&feed)
			}()

		}(feed)
	}
}
