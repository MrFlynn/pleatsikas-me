<template>
    <div id="app">
      <Navbar/>
      <div class="section">
        <div class="container">
          <Main/>
          <ContactModal/>
        </div>
      </div>
      <Footer/>
    </div>
</template>

<script>
import Main from './components/Main.vue'
import Footer from './components/Footer.vue'
import Navbar from './components/Navbar.vue'
import ContactModal from './components/ContactModal.vue'

export default {
  name: 'app',
  components: {
    Main,
    Footer,
    Navbar,
    ContactModal
  }
}
</script>

<style lang="sass">
  @import "bulma"
</style>
