<template>
  <div class="items-top relative flex min-h-[calc(100vh-3rem)] justify-center bg-gray-100 sm:items-center sm:pt-0">
    <div class="mx-auto max-w-5xl sm:px-6 lg:px-8">
      <div class="flex items-center justify-center space-x-2 pt-4">
        <div>
          <h2 class="my-5 text-center text-lg font-bold">Refund Policy</h2>
          <p>
            We have a 30-day return policy, which means you have 30 days after receiving your item to request a
            return.<br /><br />To be eligible for a return, your item must be in the same condition that you received
            it, unworn or unused, with tags, and in its original packaging. You’ll also need the receipt or proof of
            purchase. <br /><br />To start a return, you can contact us at<a
              href="mailto:info@badgerd.nl"
              target="_blank"
              rel="noopener noreferrer"
              aria-describedby="forwarding-new-window-message"
              >info@badgerd.nl</a
            >. If your return is accepted, we’ll give you an address to send the item back to us. Please be aware that
            we are not responsible for shipping costs of returned item.<br /><br />You can always contact us for any
            return question at<span>&nbsp;</span
            ><a
              href="mailto:info@badgerd.nl"
              target="_blank"
              rel="noopener noreferrer"
              aria-describedby="forwarding-new-window-message"
              >info@badgerd.nl</a
            ><span>.</span>
          </p>
          <br />
          <p>
            <strong>Damages and issues</strong> <br />Please inspect your order upon reception and contact us
            immediately if the item is defective, damaged or if you receive the wrong item, so that we can evaluate the
            issue and make it right.
          </p>
          <br />
          <p>
            <strong>Exceptions / non-returnable items</strong> <br />Certain types of items cannot be returned like
            custom products (such as special orders or personalized items). Please get in touch if you have questions or
            concerns about your specific item. <br /><br />Unfortunately, we cannot accept returns on gift cards.
          </p>
          <br />
          <p>
            <strong>Exchanges</strong> <br />The fastest way to ensure you get what you want is to return the item you
            have, and once the return is accepted, make a separate purchase for the new item.
          </p>
          <br />
          <p>
            <strong>Refunds</strong> <br />We will notify you once we’ve received and inspected your return, and let you
            know if the refund was approved or not. If approved, you’ll be automatically refunded on your original
            payment method. Please remember it can take some time for your bank or credit card company to process and
            post the refund too.
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  head() {
    return {
      link: [
        {
          hid: 'canonical',
          rel: 'canonical',
          href: `${this.$config.baseURL}${this.$route.path}`,
        },
      ],
    }
  },
}
</script>
