<template>
  <div id="app">
    <div class="columns is-gapless">
      <div class ="column is-narrow has-background-light">
        <div class="panel-box">
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
        <div class="panel-box-container">
          <div class="panel-box-centered">
            <nav class="breadcrumb has-bullet-separator" aria-label="breadcrumbs">
              <ul>
                <li><a class="subtitle is-5" href="https://github.com/MrFlynn">Github</a></li>
                <li><a class="subtitle is-5" href="https://linkedin.com/in/pleatsikas">LinkedIn</a></li>
                <li><a class="subtitle is-5">Resume</a></li>
              </ul>
            </nav>
          </div>
        </div>
        <div class="panel-footer">
          <div class="panel-box-container">
            <div class="panel-box-centered">
              <p class="panel-footer">
                Copyright Nick Pleatsikas, 2019. Source available
                <a href="https://github.com/MrFlynn/pleatsikas-me">here.</a>
              </p>
            </div>
          </div>
        </div>
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
  // This block contains customizations for Bulma.
  @import url('https://fonts.googleapis.com/css?family=Source+Sans+Pro:400,700')
  $family-sans-serif: "Source Sans Pro", sans-serif

  @import "bulma"

  // Contains customizations for site-specific elements.
  %boxing
    padding: 1rem
  
  .columns
    height: 100vh

  .panel-box
    @extend %boxing

  .panel-box-container
    @extend %boxing
    text-align: center
  
  .panel-box-centered
    display: inline-block

  .panel-footer
    position: fixed
    bottom: 0.5rem

    p
      color: $grey-light
      a
        color: $turquoise
        border-bottom: 1px solid rgba(0, 209, 178, 0)

        -webkit-transition: border 1000ms ease;
        -moz-transition: border 500ms ease;
        -ms-transition: border 500ms ease;
        -o-transition: border 500ms ease;
        transition: border 500ms ease;
        &:hover
          border-bottom: 1px solid $turquoise
    
    @media screen and (max-width: 768px)
      visibility: hidden
</style>
