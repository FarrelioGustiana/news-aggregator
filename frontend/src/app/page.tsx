"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import Image from "next/image";
import { useAuth } from "@/lib/auth/AuthContext";

export default function Home() {
  const { isAuthenticated, isLoading } = useAuth();
  const router = useRouter();

  // If authenticated, redirect to dashboard
  useEffect(() => {
    if (isAuthenticated) {
      router.push("/dashboard/articles");
    }
  }, []);

  return (
    <div className="min-h-screen bg-white">
      {/* Hero Section */}
      <div className="relative bg-gradient-to-r from-blue-600 to-blue-800 text-white py-20">
        {/* Navigation */}
        <header className="absolute top-0 left-0 right-0 z-10">
          <nav className="container mx-auto px-6 py-4 flex justify-between items-center">
            <div className="text-3xl font-bold">Pilar Credo</div>
            <div>
              {isLoading ? (
                <div className="animate-spin rounded-full h-5 w-5 border-t-2 border-b-2 border-white"></div>
              ) : isAuthenticated ? (
                <Link
                  href="/dashboard/articles"
                  className="bg-white text-blue-600 px-4 py-2 rounded-lg font-medium hover:bg-blue-50 transition-colors"
                >
                  Go to Dashboard
                </Link>
              ) : (
                <div className="space-x-4">
                  <Link
                    href="/auth/login"
                    className="text-white hover:text-blue-200 transition-colors"
                  >
                    Login
                  </Link>
                  <Link
                    href="/auth/register"
                    className="bg-white text-blue-600 px-4 py-2 rounded-lg font-medium hover:bg-blue-50 transition-colors"
                  >
                    Sign Up
                  </Link>
                </div>
              )}
            </div>
          </nav>
        </header>

        {/* Hero Content */}
        <div className="container mx-auto px-6 pt-20 flex flex-col md:flex-row items-center">
          <div className="flex flex-col w-full md:w-1/2 justify-center items-start pt-12 pb-24 px-6">
            <h1 className="font-bold text-4xl md:text-5xl leading-tight mb-4">
              Your Personalized News,
              <br /> Aggregated.
            </h1>
            <p className="text-xl mb-8">
              Get all your favorite content in one place, personalized just for
              you.
            </p>
            <div className="flex space-x-4">
              <Link
                href="/auth/register"
                className="bg-white text-blue-600 px-6 py-3 rounded-lg font-bold hover:bg-blue-50 transition-colors"
              >
                Get Started
              </Link>
              <Link
                href="#how-it-works"
                className="bg-transparent border-2 border-white text-white px-6 py-3 rounded-lg font-bold hover:bg-white hover:bg-opacity-10 transition-colors"
              >
                Learn More
              </Link>
            </div>
          </div>
          <div className="w-full md:w-1/2 py-6 text-center">
            {/* You can add a hero image here */}
            <div className="bg-white bg-opacity-10 p-8 rounded-2xl shadow-xl max-w-md mx-auto">
              <div className="space-y-4">
                {[1, 2, 3].map((item) => (
                  <div
                    key={item}
                    className="bg-white bg-opacity-20 rounded-lg p-4 flex flex-col items-start"
                  >
                    <div className="w-3/4 h-4 bg-white bg-opacity-30 rounded mb-2"></div>
                    <div className="w-1/2 h-3 bg-white bg-opacity-30 rounded"></div>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Features Section */}
      <section id="features" className="py-20 bg-gray-50">
        <div className="container mx-auto px-6">
          <h2 className="text-3xl font-bold text-center text-gray-800 mb-12">
            Why Choose Pilar Credo
          </h2>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div className="bg-white p-8 rounded-xl shadow-md">
              <div className="text-blue-600 mb-4">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-12 w-12"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9M7 16h6M7 8h6v4H7V8z"
                  />
                </svg>
              </div>
              <h3 className="text-xl font-semibold mb-2 text-gray-800">
                All Your News in One Place
              </h3>
              <p className="text-gray-600">
                Say goodbye to jumping between dozens of websites. Get all your
                news from your favorite sources in a single, unified feed.
              </p>
            </div>

            <div className="bg-white p-8 rounded-xl shadow-md">
              <div className="text-blue-600 mb-4">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-12 w-12"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z"
                  />
                </svg>
              </div>
              <h3 className="text-xl font-semibold mb-2 text-gray-800">
                Personalized Experience
              </h3>
              <p className="text-gray-600">
                Curate your own news experience by subscribing only to the
                sources and topics that interest you most.
              </p>
            </div>

            <div className="bg-white p-8 rounded-xl shadow-md">
              <div className="text-blue-600 mb-4">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-12 w-12"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
                  />
                </svg>
              </div>
              <h3 className="text-xl font-semibold mb-2 text-gray-800">
                Stay Updated Automatically
              </h3>
              <p className="text-gray-600">
                Our system automatically fetches the latest content from your
                subscribed sources, so you never miss an update.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* How It Works Section */}
      <section id="how-it-works" className="py-20">
        <div className="container mx-auto px-6">
          <h2 className="text-3xl font-bold text-center text-gray-800 mb-12">
            How It Works
          </h2>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div className="text-center">
              <div className="bg-blue-100 rounded-full w-16 h-16 flex items-center justify-center mx-auto mb-4">
                <span className="text-blue-600 text-2xl font-bold">1</span>
              </div>
              <h3 className="text-xl font-semibold mb-2 text-gray-800">
                Create an Account
              </h3>
              <p className="text-gray-600">
                Sign up for a free account to get started with your personalized
                news experience.
              </p>
            </div>

            <div className="text-center">
              <div className="bg-blue-100 rounded-full w-16 h-16 flex items-center justify-center mx-auto mb-4">
                <span className="text-blue-600 text-2xl font-bold">2</span>
              </div>
              <h3 className="text-xl font-semibold mb-2 text-gray-800">
                Subscribe to Feeds
              </h3>
              <p className="text-gray-600">
                Add your favorite news sources, blogs, and websites by
                subscribing to their RSS/Atom feeds.
              </p>
            </div>

            <div className="text-center">
              <div className="bg-blue-100 rounded-full w-16 h-16 flex items-center justify-center mx-auto mb-4">
                <span className="text-blue-600 text-2xl font-bold">3</span>
              </div>
              <h3 className="text-xl font-semibold mb-2 text-gray-800">
                Enjoy Your News Feed
              </h3>
              <p className="text-gray-600">
                Browse your personalized feed of articles from all your
                subscribed sources in one clean interface.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Call to Action */}
      <section className="py-16 bg-blue-600 text-white">
        <div className="container mx-auto px-6 text-center">
          <h2 className="text-3xl font-bold mb-4">
            Ready to Simplify Your News Experience?
          </h2>
          <p className="text-xl mb-8">
            Join Pilar Credo today and never miss important updates from your
            favorite sources.
          </p>
          <Link
            href="/auth/register"
            className="bg-white text-blue-600 px-8 py-3 rounded-lg font-bold text-lg hover:bg-blue-50 transition-colors inline-block"
          >
            Get Started Now
          </Link>
        </div>
      </section>

      {/* Footer */}
      <footer className="py-10 bg-gray-800 text-white">
        <div className="container mx-auto px-6">
          <div className="flex flex-col md:flex-row justify-between items-center">
            <div className="mb-6 md:mb-0">
              <h2 className="text-2xl font-bold">Pilar Credo</h2>
              <p className="text-gray-400">
                Your Personalized News, Aggregated.
              </p>
            </div>
            <div className="flex space-x-6">
              <Link
                href="/auth/login"
                className="hover:text-blue-400 transition-colors"
              >
                Login
              </Link>
              <Link
                href="/auth/register"
                className="hover:text-blue-400 transition-colors"
              >
                Register
              </Link>
            </div>
          </div>
          <div className="border-t border-gray-700 mt-8 pt-8 text-center text-gray-400">
            <p>
              &copy; {new Date().getFullYear()} Pilar Credo. All rights
              reserved.
            </p>
          </div>
        </div>
      </footer>
    </div>
  );
}
