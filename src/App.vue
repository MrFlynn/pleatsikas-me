<template>
    <div id="app">
      <Navbar/>
      <div class="section">
        <div class="container">
          <Home/>
          <ContactModal/>
        </div>
      </div>
      <Footer/>
    </div>
</template>

<script>
import Home from './components/Home.vue'
import Footer from './components/Footer.vue'
import Navbar from './components/Navbar.vue'
import ContactModal from './components/ContactModal.vue'

export default {
  name: 'app',
  components: {
    Home,
    Footer,
    Navbar,
    ContactModal
  }
}
</script>

<style lang="sass">
  @import "bulma"
</style>

<style>
  .footer {
    position: fixed;
    width: 100%;
    bottom: 0px;
  }
</style>
