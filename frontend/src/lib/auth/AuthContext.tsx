'use client';

import React, { createContext, useContext, useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { authAPI, userAPI } from '../api';
import { User, AuthState, UserCredentials } from '@/types';

interface AuthContextType extends AuthState {
  login: (credentials: UserCredentials) => Promise<void>;
  register: (credentials: UserCredentials) => Promise<void>;
  logout: () => Promise<void>;
  clearError: () => void;
}

// Create context with default values
const AuthContext = createContext<AuthContextType>({
  user: null,
  token: null,
  isAuthenticated: false,
  isLoading: true,
  error: null,
  login: async () => {},
  register: async () => {},
  logout: async () => {},
  clearError: () => {},
});

// Custom hook to use the auth context
export const useAuth = () => useContext(AuthContext);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ 
  children 
}) => {
  const [authState, setAuthState] = useState<AuthState>({
    user: null,
    token: null,
    isAuthenticated: false,
    isLoading: true,
    error: null,
  });
  
  const router = useRouter();

  // Check for existing authentication on component mount
  useEffect(() => {
    const checkAuth = async () => {
      try {
        // Check if we have a token stored
        const token = localStorage.getItem('token');
        
        if (!token) {
          setAuthState(prev => ({ ...prev, isLoading: false }));
          return;
        }
        
        // Fetch user profile to validate token
        const response = await userAPI.getProfile();
        const userData = response.data;
        
        setAuthState({
          user: userData,
          token,
          isAuthenticated: true,
          isLoading: false,
          error: null,
        });
      } catch (error) {
        // Token invalid or expired
        localStorage.removeItem('token');
        
        setAuthState({
          user: null,
          token: null,
          isAuthenticated: false,
          isLoading: false,
          error: null,
        });
      }
    };
    
    checkAuth();
  }, []);

  // Login function
  const login = async (credentials: UserCredentials) => {
    setAuthState(prev => ({ ...prev, isLoading: true, error: null }));
    
    try {
      const response = await authAPI.login(credentials.username, credentials.password);
      const { token } = response.data;
      
      // Fetch user data after login
      const userResponse = await userAPI.getProfile();
      const userData = userResponse.data;
      
      setAuthState({
        user: userData,
        token,
        isAuthenticated: true,
        isLoading: false,
        error: null,
      });
      
      // Redirect to dashboard
      router.push('/dashboard/articles');
    } catch (error: any) {
      const errorMsg = error.response?.data?.message || 'Failed to login. Please check your credentials.';
      
      setAuthState(prev => ({
        ...prev,
        isLoading: false,
        error: errorMsg,
      }));
    }
  };

  // Register function
  const register = async (credentials: UserCredentials) => {
    setAuthState(prev => ({ ...prev, isLoading: true, error: null }));
    
    try {
      await authAPI.register(credentials.username, credentials.password);
      
      // Login after successful registration
      await login(credentials);
    } catch (error: any) {
      const errorMsg = error.response?.data?.message || 'Registration failed. Please try again.';
      
      setAuthState(prev => ({
        ...prev,
        isLoading: false,
        error: errorMsg,
      }));
    }
  };

  // Logout function
  const logout = async () => {
    setAuthState(prev => ({ ...prev, isLoading: true }));
    
    try {
      await authAPI.logout();
      
      setAuthState({
        user: null,
        token: null,
        isAuthenticated: false,
        isLoading: false,
        error: null,
      });
      
      // Redirect to login page
      router.push('/auth/login');
    } catch (error) {
      setAuthState(prev => ({ ...prev, isLoading: false }));
    }
  };

  // Clear error
  const clearError = () => {
    setAuthState(prev => ({ ...prev, error: null }));
  };

  // Context value
  const value: AuthContextType = {
    ...authState,
    login,
    register,
    logout,
    clearError,
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthContext;
