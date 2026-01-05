
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useCartStore = defineStore('cart', () => {
  // State
  const items = ref([])

  // Load cart from localStorage on initialization
  const loadCart = () => {
    const savedCart = localStorage.getItem('jumia_cart')
    if (savedCart) {
      try {
        items.value = JSON.parse(savedCart)
      } catch (e) {
        console.error('Error loading cart:', e)
        items.value = []
      }
    }
  }

  // Save cart to localStorage
  const saveCart = () => {
    localStorage.setItem('jumia_cart', JSON.stringify(items.value))
  }

  // Getters
  const cartCount = computed(() => {
    return items.value.reduce((total, item) => total + item.quantity, 0)
  })

  const cartTotal = computed(() => {
    return items.value.reduce((total, item) => {
      return total + (item.final_price * item.quantity)
    }, 0)
  })

  const cartItems = computed(() => items.value)

  // Actions
  const addToCart = (product, quantity = 1) => {
    const existingItem = items.value.find(item => item.id === product.id)

    if (existingItem) {
      // Update quantity if item already in cart
      existingItem.quantity += quantity
    } else {
      // Add new item to cart
      items.value.push({
        id: product.id,
        name: product.name,
        image_url: product.image_url,
        price: product.price,
        final_price: product.final_price,
        discount_percentage: product.discount_percentage,
        brand: product.brand,
        stock: product.stock,
        in_stock: product.in_stock,
        quantity: quantity
      })
    }

    saveCart()
  }

  const removeFromCart = (productId) => {
    const index = items.value.findIndex(item => item.id === productId)
    if (index > -1) {
      items.value.splice(index, 1)
      saveCart()
    }
  }

  const updateQuantity = (productId, quantity) => {
    const item = items.value.find(item => item.id === productId)
    if (item) {
      if (quantity <= 0) {
        removeFromCart(productId)
      } else {
        item.quantity = quantity
        saveCart()
      }
    }
  }

  const incrementQuantity = (productId) => {
    const item = items.value.find(item => item.id === productId)
    if (item) {
      item.quantity++
      saveCart()
    }
  }

  const decrementQuantity = (productId) => {
    const item = items.value.find(item => item.id === productId)
    if (item && item.quantity > 1) {
      item.quantity--
      saveCart()
    }
  }

  const clearCart = () => {
    items.value = []
    saveCart()
  }

  // Initialize cart from localStorage
  loadCart()

  return {
    items,
    cartCount,
    cartTotal,
    cartItems,
    addToCart,
    removeFromCart,
    updateQuantity,
    incrementQuantity,
    decrementQuantity,
    clearCart,
    loadCart
  }
})
