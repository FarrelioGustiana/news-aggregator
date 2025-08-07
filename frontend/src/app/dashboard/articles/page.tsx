"use client";

import { useState, useEffect } from "react";
import { articleAPI } from "@/lib/api";
import { Article } from "@/types";
import Link from "next/link";
import { toast } from "react-hot-toast";
import ArticleModal from "@/components/articles/ArticleModal";

export default function ArticlesPage() {
  const [articles, setArticles] = useState<Article[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [page, setPage] = useState<number>(1);
  const [hasMore, setHasMore] = useState<boolean>(true);
  const [selectedArticle, setSelectedArticle] = useState<Article | null>(null);
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
  const pageSize = 10;

  const fetchArticles = async (pageNum: number) => {
    setLoading(true);
    setError(null);
    try {
      const response = await articleAPI.getArticlesForUser(pageNum, pageSize);
      const newArticles = response.data.articles || [];

      if (pageNum === 1) {
        setArticles(newArticles);
      } else {
        setArticles((prevArticles) => [...prevArticles, ...newArticles]);
      }

      // Check if we have more articles to load
      setHasMore(newArticles.length === pageSize);
    } catch (err: any) {
      setError(err.response?.data?.message || "Failed to load articles");
      toast.error("Failed to load articles");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchArticles(page);
  }, [page]);

  const loadMoreArticles = () => {
    if (!loading && hasMore) {
      setPage((prevPage) => prevPage + 1);
    }
  };

  const handleRefresh = () => {
    setPage(1);
    fetchArticles(1);
    toast.success("Refreshing your feed");
  };

  const openArticleModal = (article: Article) => {
    setSelectedArticle(article);
    setIsModalOpen(true);
  };

  const closeArticleModal = () => {
    setIsModalOpen(false);
  };

  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleDateString("en-US", {
      year: "numeric",
      month: "short",
      day: "numeric",
    });
  };

  return (
    <div className="p-6">
      <div className="flex justify-between items-center mb-6">
        <div>
          <h1 className="text-3xl font-bold text-gray-800">Your News Feed</h1>
          <p className="text-gray-600">
            Latest articles from your subscribed sources
          </p>
        </div>
        <button
          onClick={handleRefresh}
          disabled={loading && page === 1}
          className="flex items-center bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg transition-colors disabled:bg-blue-400"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="h-5 w-5 mr-2"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"
            />
          </svg>
          Refresh
        </button>
      </div>

      {/* Filters - for future implementation */}
      <div className="mb-6 bg-gray-50 p-3 rounded-lg border border-gray-200">
        <div className="flex flex-wrap gap-2">
          <span className="text-sm text-gray-600 mr-2">Quick filters:</span>
          <button className="bg-white border border-gray-200 hover:bg-gray-50 text-gray-700 text-sm px-3 py-1 rounded-full">
            All sources
          </button>
          <button className="bg-white border border-gray-200 hover:bg-gray-50 text-gray-700 text-sm px-3 py-1 rounded-full">
            Today
          </button>
          <button className="bg-white border border-gray-200 hover:bg-gray-50 text-gray-700 text-sm px-3 py-1 rounded-full">
            This week
          </button>
        </div>
      </div>

      {loading && page === 1 && (
        <div className="flex justify-center my-10">
          <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-600"></div>
        </div>
      )}

      {error && (
        <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded mb-6">
          {error}
        </div>
      )}

      {!loading && articles.length === 0 && !error && (
        <div className="text-center py-10 bg-gray-50 rounded-lg border border-gray-100">
          <h2 className="text-xl font-semibold mb-2">No articles found</h2>
          <p className="text-gray-600 mb-4">
            You don&apos;t have any articles in your feed yet.
          </p>
          <Link
            href="/dashboard/feeds"
            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded transition-colors"
          >
            Browse Available Feeds
          </Link>
        </div>
      )}

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        {articles.map((article) => (
          <div
            key={article.id}
            className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow cursor-pointer"
            onClick={() => openArticleModal(article)}
          >
            <div className="p-6">
              <h2 className="text-xl font-bold mb-2 hover:text-blue-600">
                {article.title}
              </h2>
              <div className="flex items-center text-gray-500 text-sm mb-3">
                <span className="mr-2">{formatDate(article.published_at)}</span>
                <span className="bg-blue-100 text-blue-800 text-xs px-2 py-1 rounded-full">
                  {article.feed_name || "News"}
                </span>
              </div>
              <p className="text-gray-600 mb-4 line-clamp-3">
                {article.description?.replace(/<[^>]*>/g, "") ||
                  "No description available"}
              </p>
              <button
                onClick={(e) => {
                  e.stopPropagation();
                  openArticleModal(article);
                }}
                className="text-blue-600 hover:text-blue-800 font-medium flex items-center"
              >
                Read more
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-4 w-4 ml-1"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M14 5l7 7m0 0l-7 7m7-7H3"
                  />
                </svg>
              </button>
            </div>
          </div>
        ))}
      </div>

      {articles.length > 0 && (
        <div className="mt-8 text-center">
          {loading && page > 1 && (
            <div className="flex justify-center my-4">
              <div className="animate-spin rounded-full h-8 w-8 border-t-2 border-b-2 border-blue-600"></div>
            </div>
          )}

          {hasMore && !loading && (
            <button
              onClick={loadMoreArticles}
              className="bg-white border border-gray-300 hover:bg-gray-50 text-gray-700 font-medium py-2 px-6 rounded-md shadow-sm"
            >
              Load More Articles
            </button>
          )}

          {!hasMore && articles.length > 0 && (
            <p className="text-gray-500 mt-4">
              You've reached the end of your feed.
            </p>
          )}
        </div>
      )}

      {selectedArticle && (
        <ArticleModal
          article={selectedArticle}
          isOpen={isModalOpen}
          onClose={closeArticleModal}
        />
      )}
    </div>
  );
}
