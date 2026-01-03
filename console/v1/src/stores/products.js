import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

const API_BASE = 'http://localhost:8080/api/v1'

export const useProductStore = defineStore('products', () => {
  // State
  const products = ref([])
  const searchResults = ref([])
  const currentProduct = ref(null)
  const loading = ref(false)
  const error = ref(null)
  const searchQuery = ref('')
  const currentPage = ref(1)
  const pageSize = ref(20)
  const totalProducts = ref(0)

  // Computed
  const hasProducts = computed(() => products.value.length > 0)
  const hasSearchResults = computed(() => searchResults.value.length > 0)
  const totalPages = computed(() => Math.ceil(totalProducts.value / pageSize.value))

  // Actions
  async function fetchProducts(page = 1, size = 20) {
    loading.value = true
    error.value = null

    try {
      const response = await fetch(`${API_BASE}/products?page=${page}&page_size=${size}`)

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const data = await response.json()
      products.value = data.products || []
      totalProducts.value = data.total || 0
      currentPage.value = page

      return data
    } catch (err) {
      error.value = err.message
      console.error('Error fetching products:', err)
      return null
    } finally {
      loading.value = false
    }
  }

  async function searchProducts(query, page = 1, size = 20) {
    if (!query || query.trim() === '') {
      searchResults.value = []
      searchQuery.value = ''
      return
    }

    loading.value = true
    error.value = null
    searchQuery.value = query

    try {
      const response = await fetch(
        `${API_BASE}/products/search?q=${encodeURIComponent(query)}&page=${page}&page_size=${size}`
      )

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const data = await response.json()
      searchResults.value = data.products || []
      totalProducts.value = data.total || 0
      currentPage.value = page

      return data
    } catch (err) {
      error.value = err.message
      console.error('Error searching products:', err)
      searchResults.value = []
      return null
    } finally {
      loading.value = false
    }
  }

  async function getProductsByCategory(category, page = 1, size = 20) {
    loading.value = true
    error.value = null

    try {
      const response = await fetch(
        `${API_BASE}/products/category/${encodeURIComponent(category)}?page=${page}&page_size=${size}`
      )

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const data = await response.json()
      products.value = data.products || []
      totalProducts.value = data.total || 0
      currentPage.value = page

      return data
    } catch (err) {
      error.value = err.message
      console.error('Error fetching products by category:', err)
      return null
    } finally {
      loading.value = false
    }
  }

  async function getProduct(id) {
    loading.value = true
    error.value = null

    try {
      const response = await fetch(`${API_BASE}/products/${id}`)

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const data = await response.json()
      currentProduct.value = data.product || null

      return data
    } catch (err) {
      error.value = err.message
      console.error('Error fetching product:', err)
      currentProduct.value = null
      return null
    } finally {
      loading.value = false
    }
  }

  function clearSearch() {
    searchResults.value = []
    searchQuery.value = ''
    error.value = null
  }

  function clearError() {
    error.value = null
  }

  return {
    // State
    products,
    searchResults,
    currentProduct,
    loading,
    error,
    searchQuery,
    currentPage,
    pageSize,
    totalProducts,

    // Computed
    hasProducts,
    hasSearchResults,
    totalPages,

    // Actions
    fetchProducts,
    searchProducts,
    getProductsByCategory,
    getProduct,
    clearSearch,
    clearError,
  }
})
