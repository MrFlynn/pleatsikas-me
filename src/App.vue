<template>
  <div id="app">
    <div class="container is-max-widescreen">
      <div class="columns is-gapless">
        <div class ="column is-narrow has-background-light">
          <div class="panel-box">
            <div class="person">
              <img class="profile-image" src="./assets/profile.jpg" alt="profile.jpg" width="192" height="192">
            </div>
            <div class="person content">
              <h1 class="title is-2">Nick Pleatsikas</h1>
              <h3 class="subtitle">Software Engineer</h3>
              <button @click="showModal()" class="button is-info">Contact Me</button>
            </div>
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
                    <a class="subtitle is-5" href="resume.pdf" style="padding-right: 0">
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
                  © 2020 Nick Pleatsikas. Source available
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
  @import url('https://fonts.googleapis.com/css?family=Work+Sans:400,700')

  // Bulma variable overiddes.
  $family-sans-serif: "Work Sans", sans-serif
  $tablet: 900px
  
  // Custom color scheme.
  $dark: #010101
  $light: #EFEFEF
  $cyan: #8EB8B6
  $blue: #3E727F
  $green: $blue

  // Post imports
  @import "bulma"

  // Contains customizations for site-specific elements.
  .content
    p, li
      color: $dark
      font-size: 1.3rem

  .person
    text-align: center
    margin: 0.5rem auto 0.5rem auto

  .profile-image
    border-radius: 50%

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

  .tag
    margin-right: 0.4rem
    @media screen and (max-width: $tablet)
      margin-bottom: 0.4rem

    &.link
      a, p
        color: white
      
      &:hover
        background-color: $blue

        transition: 500ms ease
        -webkit-transition: 500ms ease
        -moz-transition: 500ms ease
        -ms-transition: 500ms ease
        -o-transition: 500ms ease
</style>
