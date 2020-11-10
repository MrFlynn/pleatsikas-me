<template>
  <div v-if="!isHidden" class="modal is-active">
    <div class="modal-background"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <strong class="modal-card-title">Contact Form</strong>
        <button @click="isHidden = true" class="delete" aria-label="close"></button>
      </header>
      <section class="modal-card-body">
        <form name="contact" method="POST" data-netlify="true">
          <div class="field">
            <label class="label">Name</label>
            <div class="control">
              <input class="input" type="text" placeholder="Name" name="name">
            </div>
          </div>
          <div class="field">
            <label class="label">Email</label>
            <div class="control">
              <input class="input" type="text" placeholder="name@email.com" name="email">
            </div>
          </div>
          <div class="field">
            <label class="label">Message</label>
            <div class="control">
              <textarea class="textarea" placeholder="Message here" name="message_body"></textarea>
            </div>
          </div>
        </form>
      </section>
      <footer class="modal-card-foot">
        <button class="button is-success" type="submit">Submit</button>
        <button @click="isHidden = true" class="button">Cancel</button>
      </footer>
    </div>
  </div>
</template>

<script>
import {eventBus} from '../main';

export default {
  name: 'ContactModal',
  created() {
    eventBus.$on('modalHidden', (message) => {
      this.isHidden = message;
    });
  },
  props: {
    isHidden: {
      type: Boolean,
      default: true
    }
  }
}
</script>

<style lang="sass">
  =keyframes($name)
    @-webkit-keyframes #{$name}
      @content
    @-moz-keyframes #{$name}
      @content
    @-ms-keyframes #{$name}
      @content
    @keyframes #{$name}
      @content

  +keyframes(window-fade-in)
    0%
      opacity: 0
    100%
      opacity: 1

  .modal
    animation: window-fade-in 0.4s
    -moz-animation: window-fade-in 0.4s
    -webkit-animation: window-fade-in 0.4s
    -o-animation: window-fade-in 0.4s
</style>
