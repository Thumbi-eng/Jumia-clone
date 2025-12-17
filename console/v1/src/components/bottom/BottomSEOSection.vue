<template>
  <v-card
    elevation="0"
    rounded="lg"
    class="pa-6"
    style="max-width: 1550px; margin: 0 auto"
  >
    <!-- SEO H2 -->
    <h2 class="text-h6 font-weight-bold mb-4">
      Jumia Kenya – Online Shopping, Best Deals & Trusted Marketplace
    </h2>

    <p class="text-body-2 mb-6">
      Shop online in Kenya for electronics, phones, fashion, appliances,
      groceries, beauty products and more at unbeatable prices. Enjoy fast
      delivery, secure payments, and exclusive deals from trusted brands.
    </p>

    <!-- SEO Sections -->
    <h3 class="text-subtitle-1 font-weight-bold mb-2">
      Why Shop Online on Jumia Kenya?
    </h3>

    <p class="text-body-2 mb-6">
      Jumia Kenya is the leading e-commerce platform offering premium brands
      such as
      <strong>Samsung</strong>, <strong>Apple</strong>, <strong>Tecno</strong>,
      <strong>Infinix</strong>, <strong>Xiaomi</strong>, <strong>HP</strong> and
      more. Discover flash sales, stock clearance offers, Black Friday deals and
      exclusive promotions when you buy online.
    </p>

    <v-divider class="my-6" />

    <!-- FAQ Schema Container -->
    <section itemscope itemtype="https://schema.org/FAQPage">
      <h3 class="text-subtitle-1 font-weight-bold mb-4">
        Frequently Asked Questions (FAQs)
      </h3>

      <v-expansion-panels variant="accordion" elevation="0">
        <v-expansion-panel v-for="(faq, index) in faqs" :key="index">
          <v-expansion-panel-title>
            <span class="font-weight-medium text-body-2" itemprop="name">
              {{ faq.question }}
            </span>
          </v-expansion-panel-title>

          <v-expansion-panel-text>
            <div
              class="text-body-2"
              itemprop="acceptedAnswer"
              itemscope
              itemtype="https://schema.org/Answer"
            >
              <span itemprop="text">
                {{ faq.answer }}
              </span>
            </div>
          </v-expansion-panel-text>
        </v-expansion-panel>
      </v-expansion-panels>
    </section>
  </v-card>
</template>
<script setup>
const faqs = [
  {
    question: "How do I place an order on Jumia Kenya?",
    answer:
      "To place an order on Jumia Kenya, browse the website, select your preferred products, add them to your cart, and proceed to checkout. Follow the on-screen instructions to complete your purchase securely.",
  },
  {
    question: "What payment methods are available on Jumia?",
    answer:
      "Jumia accepts multiple payment methods including credit and debit cards, mobile money, and cash on delivery, allowing you to choose the most convenient option.",
  },
  {
    question: "How long does delivery take in Kenya?",
    answer:
      "Delivery times vary depending on your location and the product. Estimated delivery dates are displayed during checkout for Nairobi, Mombasa, Nakuru and other regions in Kenya.",
  },
  {
    question: "Does Jumia offer free delivery?",
    answer:
      "Yes, Jumia offers free delivery on selected products and during promotional campaigns. Items marked with free shipping qualify for this service.",
  },
  {
    question: "What is Jumia’s return and refund policy?",
    answer:
      "Jumia provides a hassle-free return and refund policy. If you are not satisfied with your order, you can initiate a return request through your account within the specified return period.",
  },
  {
    question: "Is it safe to shop online on Jumia?",
    answer:
      "Yes, shopping on Jumia is safe. The platform uses secure payment systems and protects customer data with strict security measures.",
  },
];
</script>
