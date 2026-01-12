<template>
  <div>
    <h1 class="text-h5 font-weight-regular mb-6">Dashboard</h1>

    <!-- Stats Cards -->
    <v-row class="mb-6">
      <v-col cols="12" md="3">
        <v-card elevation="0" color="primary">
          <v-card-text class="text-white">
            <div class="text-h4 font-weight-bold">1,234</div>
            <div class="text-body-2">Total Products</div>
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="12" md="3">
        <v-card elevation="0" color="success">
          <v-card-text class="text-white">
            <div class="text-h4 font-weight-bold">456</div>
            <div class="text-body-2">Total Orders</div>
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="12" md="3">
        <v-card elevation="0" color="orange">
          <v-card-text class="text-white">
            <div class="text-h4 font-weight-bold">789</div>
            <div class="text-body-2">Total Customers</div>
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="12" md="3">
        <v-card elevation="0" color="info">
          <v-card-text class="text-white">
            <div class="text-h4 font-weight-bold">KSh 2.5M</div>
            <div class="text-body-2">Total Revenue</div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <!-- Recent Orders -->
    <v-card elevation="0">
      <v-card-title class="text-subtitle-1 font-weight-medium">
        Recent Orders
      </v-card-title>
      <v-divider></v-divider>
      <v-card-text>
        <v-table>
          <thead>
            <tr>
              <th>Order ID</th>
              <th>Customer</th>
              <th>Date</th>
              <th>Amount</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>#12345</td>
              <td>John Doe</td>
              <td>Jan 12, 2026</td>
              <td>KSh 15,000</td>
              <td><v-chip color="success" size="small">Completed</v-chip></td>
            </tr>
            <tr>
              <td>#12346</td>
              <td>Jane Smith</td>
              <td>Jan 12, 2026</td>
              <td>KSh 8,500</td>
              <td><v-chip color="warning" size="small">Pending</v-chip></td>
            </tr>
          </tbody>
        </v-table>
      </v-card-text>
    </v-card>
  </div>
</template>

<script setup>
// Admin Dashboard
</script>
