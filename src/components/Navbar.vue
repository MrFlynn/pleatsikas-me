<template>
  <nav class="navbar is-spaced is-light" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
      <figure class="image is-48x48">
        <img class="is-rounded" src="../assets/mini.png">
      </figure>
    </div>
    <div class="navbar-menu">
      <div class="navbar-start">
        <p class="navbar-item is-3">Nick Pleatsikas</p>
      </div>
      <div class="navbar-end">
        <button @click="showModal()" class="button is-info">Contact Me</button>
      </div>
    </div>
  </nav>
</template>

<script>
import {eventBus} from '../main';

export default {
  name: 'Navbar',
  methods: {
    showModal() {
      eventBus.$emit('modalHidden', false)
    }
  }
}
</script>
