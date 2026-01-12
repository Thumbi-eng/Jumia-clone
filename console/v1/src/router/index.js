import { createRouter, createWebHistory } from 'vue-router'

// Layouts
import CustomerLayout from '@/layouts/CustomerLayout.vue'
import AdminLayout from '@/layouts/AdminLayout.vue'

// Pages (route-level components)
import Home from '@/views/Home.vue'
import Cart from '@/views/Cart.vue'
import Checkout from '@/views/Checkout.vue'
import SearchResults from '@/views/SearchResults.vue'
import Login from '@/views/Login.vue'
import Register from '@/views/Register.vue'
import ProductDetail from '@/views/ProductDetail.vue'

// Admin pages
import Dashboard from '@/views/admin/Dashboard.vue'
import AddProducts from '@/views/admin/AddProducts.vue'
import ProductsList from '@/views/admin/ProductsList.vue'
import FlashSales from '@/views/admin/FlashSales.vue'
import TopDeals from '@/views/admin/TopDeals.vue'
import AdminOrders from '@/views/admin/Orders.vue'
import Customers from '@/views/admin/Customers.vue'
import Categories from '@/views/admin/Categories.vue'
import Analytics from '@/views/admin/Analytics.vue'
import Settings from '@/views/admin/Settings.vue'

// Customer pages
import AccountOverview from '@/views/customer/AccountOverview.vue'
import Orders from '@/views/customer/Orders.vue'
import Inbox from '@/views/customer/Inbox.vue'
import PendingReviews from '@/views/customer/PendingReviews.vue'
import Vouchers from '@/views/customer/Vouchers.vue'
import Wishlist from '@/views/customer/Wishlist.vue'
import FollowedSellers from '@/views/customer/FollowedSellers.vue'
import RecentlyViewed from '@/views/customer/RecentlyViewed.vue'
import AccountManagement from '@/views/customer/AccountManagement.vue'
import PaymentSettings from '@/views/customer/PaymentSettings.vue'
import AddressBook from '@/views/customer/AddressBook.vue'
import NewsletterPreferences from '@/views/customer/NewsletterPreferences.vue'
import CloseAccount from '@/views/customer/CloseAccount.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'home', component: Home },
    { path: '/search', name: 'search', component: SearchResults },
    { path: '/product/:id', name: 'product-detail', component: ProductDetail },
    { path: '/login', name: 'login', component: Login },
    { path: '/register', name: 'register', component: Register },

    // Admin routes with layout
    {
      path: '/admin',
      component: AdminLayout,
      children: [
        {
          path: 'dashboard',
          name: 'admin-dashboard',
          component: Dashboard,
        },
        {
          path: 'products',
          name: 'admin-products-list',
          component: ProductsList,
        },
        {
          path: 'products/add',
          name: 'admin-products-add',
          component: AddProducts,
        },
        {
          path: 'flash-sales',
          name: 'admin-flash-sales',
          component: FlashSales,
        },
        {
          path: 'top-deals',
          name: 'admin-top-deals',
          component: TopDeals,
        },
        {
          path: 'orders',
          name: 'admin-orders',
          component: AdminOrders,
        },
        {
          path: 'customers',
          name: 'admin-customers',
          component: Customers,
        },
        {
          path: 'categories',
          name: 'admin-categories',
          component: Categories,
        },
        {
          path: 'analytics',
          name: 'admin-analytics',
          component: Analytics,
        },
        {
          path: 'settings',
          name: 'admin-settings',
          component: Settings,
        },
      ],
    },

    { path: '/cart', name: 'cart', component: Cart },
    { path: '/checkout/summary', name: 'checkout', component: Checkout },
    {
      path: '/customer',
      component: CustomerLayout,
      children: [
        {
          path: 'account/index',
          name: 'customer-account',
          component: AccountOverview,
        },
        {
          path: 'order/index',
          name: 'customer-orders',
          component: Orders,
        },
        {
          path: 'inbox/index',
          name: 'customer-inbox',
          component: Inbox,
        },
        {
          path: 'reviews/index',
          name: 'customer-reviews',
          component: PendingReviews,
        },
        {
          path: 'vouchers/index',
          name: 'customer-vouchers',
          component: Vouchers,
        },
        {
          path: 'wishlist/index',
          name: 'customer-wishlist',
          component: Wishlist,
        },
        {
          path: 'followed-sellers/index',
          name: 'customer-followed-sellers',
          component: FollowedSellers,
        },
        {
          path: 'recently-viewed/index',
          name: 'customer-recently-viewed',
          component: RecentlyViewed,
        },
        {
          path: 'account-management/index',
          name: 'customer-account-management',
          component: AccountManagement,
        },
        {
          path: 'payment-settings/index',
          name: 'customer-payment-settings',
          component: PaymentSettings,
        },
        {
          path: 'address-book/index',
          name: 'customer-address-book',
          component: AddressBook,
        },
        {
          path: 'newsletter/index',
          name: 'customer-newsletter',
          component: NewsletterPreferences,
        },
        {
          path: 'close-account/index',
          name: 'customer-close-account',
          component: CloseAccount,
        },
      ],
    },
  ],
})

export default router
