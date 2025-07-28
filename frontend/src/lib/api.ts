import axios from 'axios';

// Create an Axios instance with custom configuration
const api = axios.create({
  baseURL: 'http://localhost:8080', // Backend API URL
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true, // Important for cookies/authentication
});

// Request interceptor for adding auth token to requests
api.interceptors.request.use(
  (config) => {
    // Get token from localStorage or cookie if needed
    const token = localStorage.getItem('token');
    
    // If token exists, add it to the headers
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    
    return config;
  },
  (error) => Promise.reject(error)
);

// Response interceptor for handling errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    const { status } = error.response || {};
    
    // Handle authentication errors
    if (status === 401) {
      // Clear authentication data
      localStorage.removeItem('token');
      
      // Redirect to login page if not already there
      if (window.location.pathname !== '/auth/login') {
        window.location.href = '/auth/login';
      }
    }
    
    return Promise.reject(error);
  }
);

// Authentication API calls
export const authAPI = {
  register: async (username: string, password: string) => {
    return api.post('/api/auth/register', { username, password });
  },
  
  login: async (username: string, password: string) => {
    const response = await api.post('/api/auth/login', { username, password });
    // Store token in localStorage upon successful login
    if (response.data && response.data.token) {
      localStorage.setItem('token', response.data.token);
    }
    return response;
  },
  
  logout: async () => {
    // Clear local token
    localStorage.removeItem('token');
    // No need for API call as the backend is stateless with JWT
  },
};

// User API calls
export const userAPI = {
  getProfile: async () => {
    return api.get('/api/users/me');
  },
  
  updateProfile: async (userData: { username?: string; password?: string }) => {
    return api.put('/api/users/me', userData);
  },
};

// Feed API calls
export const feedAPI = {
  getAllFeeds: async () => {
    return api.get('/api/feeds');
  },
  
  getFeedById: async (id: number) => {
    return api.get(`/api/feeds/${id}`);
  },
  
  createFeed: async (feedData: { name: string; url: string }) => {
    return api.post('/api/feeds', feedData);
  },
  
  updateFeed: async (id: number, feedData: { name?: string; url?: string }) => {
    return api.put(`/api/feeds/${id}`, feedData);
  },
  
  deleteFeed: async (id: number) => {
    return api.delete(`/api/feeds/${id}`);
  },
};

// Subscription API calls
export const subscriptionAPI = {
  getUserSubscriptions: async () => {
    return api.get('/api/subscriptions');
  },
  
  subscribeToFeed: async (feedId: number) => {
    return api.post('/api/subscriptions', { feed_id: feedId });
  },
  
  unsubscribeFromFeed: async (feedId: number) => {
    return api.delete(`/api/subscriptions/${feedId}`);
  },
  
  checkSubscriptionStatus: async (feedId: number) => {
    return api.get(`/api/subscriptions/${feedId}/status`);
  },
};

// Article API calls
export const articleAPI = {
  getArticlesForUser: async (page: number = 1, pageSize: number = 10) => {
    return api.get(`/api/articles?page=${page}&pageSize=${pageSize}`);
  },
  
  getArticleById: async (id: number) => {
    return api.get(`/api/articles/${id}`);
  },
};

export default api;
