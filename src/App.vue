<template>
  <div id="app">
    <div class="columns is-gapless">
      <div class ="column is-narrow has-background-light">
        <div class="panel-box">
          <article class="media">
            <div class="media-left">
              <figure class="image is-128x128">
                <img class="is-rounded" src="./assets/profile.jpg" alt="profile.jpg">
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
                <li>
                  <a class="subtitle is-5" href="https://github.com/MrFlynn">
                    <p class="underline">Github</p>
                  </a>
                </li>
                <li>
                  <a class="subtitle is-5" href="https://linkedin.com/in/pleatsikas">
                    <p class="underline">LinkedIn</p>
                  </a>
                </li>
                <li>
                  <a class="subtitle is-5">
                    <p class="underline">Resume</p>
                  </a>
                </li>
              </ul>
            </nav>
          </div>
        </div>
        <div class="panel-footer">
          <div class="panel-box-container">
            <div class="panel-box-centered">
              <p class="panel-footer">
                © 2019 Nick Pleatsikas. Source available
                <a href="https://github.com/MrFlynn/pleatsikas-me">here.</a>
              </p>
            </div>
          </div>
        </div>
      </div>
      <div class="column">
        <div class="panel-box">
          <AboutMe/>
          <hr/>
          <WorkExperience/>
          <hr style="margin-top: 0"/>
          <Projects/>
        </div>
      </div>
    </div>
    <ContactModal/>
  </div>
</template>

<script>
import {eventBus} from './main';

import ContactModal from './components/ContactModal.vue'
import AboutMe from './components/AboutMe.vue'
import WorkExperience from './components/WorkExperience.vue'
import Projects from './components/Projects.vue'

export default {
  name: 'app',
  components: {
    ContactModal,
    AboutMe,
    WorkExperience,
    Projects
  },
  methods: {
    showModal() {
      eventBus.$emit('modalHidden', false)
    }
  }
}
</script>

<style lang="sass">
  // Imports
  @import "./mixins.sass"
  @import url('https://fonts.googleapis.com/css?family=Source+Sans+Pro:400,700')

  // Bulma variable overiddes.
  $family-sans-serif: "Source Sans Pro", sans-serif
  $tablet: 900px
  
  // Custom color scheme.
  $dark: #13293D
  $light: #EFEFEF
  $cyan: #2A9D8F
  $blue: #070707
  $green: #2A9D8F

  // Post imports
  @import "bulma"

  // Contains customizations for site-specific elements.
  %boxing
    padding: 1rem
  
  .columns
    min-height: 100vh

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
        +fade-in-underline($turquoise)
        color: $turquoise
    
    @media screen and (max-width: $tablet)
      visibility: hidden

  .underline
    +fade-in-underline($blue)
</style>
