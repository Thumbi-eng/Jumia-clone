import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

const API_BASE = 'http://localhost:8080/api/v1'

export const useAuthStore = defineStore('auth', () => {
  // State
  const user = ref(null)
  const accessToken = ref(localStorage.getItem('access_token') || null)
  const refreshToken = ref(localStorage.getItem('refresh_token') || null)
  const loading = ref(false)
  const error = ref(null)

  // Computed
  const isAuthenticated = computed(() => !!accessToken.value && !!user.value)
  const isAdmin = computed(() => {
    if (!user.value) return false
    // Check if user has admin role from backend
    if (user.value.role === 'admin' || user.value.is_admin === true) {
      return true
    }
    // Temporary: Hardcoded admin email for development
    if (user.value.email === 'mustangmerchants@gmail.com') {
      return true
    }
    return false
  })
  const fullName = computed(() => {
    if (!user.value) return ''
    return `${user.value.first_name} ${user.value.last_name}`
  })

  // Actions
  async function register(userData) {
    loading.value = true
    error.value = null

    try {
      const response = await fetch(`${API_BASE}/users/register`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(userData),
      })

      if (!response.ok) {
        const data = await response.json()
        throw new Error(data.error || data.message || 'Registration failed')
      }

      const data = await response.json()

      // Set tokens (backend returns 'token' instead of 'access_token')
      accessToken.value = data.token || data.access_token
      refreshToken.value = data.refresh_token

      // Store tokens in localStorage
      localStorage.setItem('access_token', accessToken.value)
      localStorage.setItem('refresh_token', data.refresh_token)

      // Set user data
      user.value = data.user

      return data
    } catch (err) {
      error.value = err.message
      console.error('Registration error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function login(email, password) {
    loading.value = true
    error.value = null

    try {
      const response = await fetch(`${API_BASE}/users/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      })

      if (!response.ok) {
        const data = await response.json()
        throw new Error(data.error || data.message || 'Login failed')
      }

      const data = await response.json()

      // Set tokens (backend returns 'token' instead of 'access_token')
      accessToken.value = data.token || data.access_token
      refreshToken.value = data.refresh_token

      // Store tokens in localStorage
      localStorage.setItem('access_token', accessToken.value)
      localStorage.setItem('refresh_token', data.refresh_token)

      // Set user data
      user.value = data.user

      return data
    } catch (err) {
      error.value = err.message
      console.error('Login error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function fetchUser() {
    if (!accessToken.value) {
      return null
    }

    loading.value = true
    error.value = null

    try {
      const response = await fetch(`${API_BASE}/users/me`, {
        headers: {
          'Authorization': `Bearer ${accessToken.value}`,
        },
      })

      if (!response.ok) {
        if (response.status === 401) {
          // Token expired, try to refresh
          await refreshAccessToken()
          return await fetchUser()
        }
        throw new Error('Failed to fetch user')
      }

      const data = await response.json()
      user.value = data.user

      return data.user
    } catch (err) {
      error.value = err.message
      console.error('Fetch user error:', err)
      // If fetch fails, clear auth
      logout()
      return null
    } finally {
      loading.value = false
    }
  }

  async function updateUser(userData) {
    if (!accessToken.value || !user.value) {
      throw new Error('Not authenticated')
    }

    loading.value = true
    error.value = null

    try {
      const response = await fetch(`${API_BASE}/users/${user.value.id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${accessToken.value}`,
        },
        body: JSON.stringify(userData),
      })

      if (!response.ok) {
        const data = await response.json()
        throw new Error(data.error || data.message || 'Update failed')
      }

      const data = await response.json()
      user.value = data.user

      return data.user
    } catch (err) {
      error.value = err.message
      console.error('Update user error:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function refreshAccessToken() {
    if (!refreshToken.value) {
      logout()
      throw new Error('No refresh token')
    }

    try {
      const response = await fetch(`${API_BASE}/users/refresh`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ refresh_token: refreshToken.value }),
      })

      if (!response.ok) {
        logout()
        throw new Error('Token refresh failed')
      }

      const data = await response.json()

      // Update access token
      accessToken.value = data.access_token
      localStorage.setItem('access_token', data.access_token)

      // Update refresh token if provided
      if (data.refresh_token) {
        refreshToken.value = data.refresh_token
        localStorage.setItem('refresh_token', data.refresh_token)
      }

      return data.access_token
    } catch (err) {
      console.error('Refresh token error:', err)
      logout()
      throw err
    }
  }

  function logout() {
    // Clear state
    user.value = null
    accessToken.value = null
    refreshToken.value = null
    error.value = null

    // Clear localStorage
    localStorage.removeItem('access_token')
    localStorage.removeItem('refresh_token')
  }

  function clearError() {
    error.value = null
  }

  // Initialize auth state on store creation
  if (accessToken.value) {
    fetchUser()
  }

  return {
    // State
    user,
    accessToken,
    refreshToken,
    loading,
    error,

    // Computed
    isAuthenticated,
    isAdmin,
    fullName,

    // Actions
    register,
    login,
    logout,
    fetchUser,
    updateUser,
    refreshAccessToken,
    clearError,
  }
})
