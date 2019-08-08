<template>
  <div id="app">
    <div class="columns is-gapless">
      <div class ="column is-one-quarter has-background-light">
        <article class="media">
          <div class="media-left">
            <figure class="image is-128x128">
              <img class="is-rounded" src="./assets/profile.jpg">
            </figure>
          </div>
          <div class="media-content">
            <div class="content">
              <h1 class="title">Nick Pleatsikas</h1>
              <h3 class="subtitle">Software Engineer</h3>
              <button @click="showModal()" class="button is-info">Contact Me</button>
            </div>
          </div>
        </article>
      </div>
      <div class="column">
        <h1>This is the content region</h1>
      </div>
    </div>
    <ContactModal/>
  </div>
</template>

<script>
import {eventBus} from './main';

import ContactModal from './components/ContactModal.vue'

export default {
  name: 'app',
  components: {
    ContactModal
  },
  methods: {
    showModal() {
      eventBus.$emit('modalHidden', false)
    }
  }
}
</script>

<style lang="sass">
  @import "bulma"
</style>

<style>
  .columns {
    height: 100vh;
  }
</style>
