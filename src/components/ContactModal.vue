<template>
  <div v-if="!isHidden" class="modal is-active">
    <div class="modal-background"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <strong class="modal-card-title">Contact me</strong>
        <button @click="isHidden = true" class="delete" aria-label="close"></button>
      </header>
      <section class="modal-card-body">
        <form name="contact" method="POST" @submit.prevent="submitForm" data-netlify="true">
          <div class="field">
            <label class="label">Name</label>
            <div class="control">
              <input v-model="form.name" class="input" type="text" placeholder="Name" name="name">
            </div>
          </div>
          <div class="field">
            <label class="label">Email</label>
            <div class="control">
              <input v-model="form.email" class="input" type="text" placeholder="name@email.com" inputmode="email" name="email">
            </div>
          </div>
          <div class="field">
            <label class="label">Message</label>
            <div class="control">
              <textarea v-model="form.message_body" class="textarea" placeholder="Message here" name="message_body"></textarea>
            </div>
          </div>
          <button class="button is-success" type="submit">Submit</button>
        </form>
      </section>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

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
  },
  data() {
    return {
      form: {
        name: "",
        email: "",
        message_body: ""
      }
    }
  },
  methods: {
    createNotification(type, message) {
      eventBus.$emit('notificationHidden', {
        type,
        message
      });
    },
    encode(data) {
      return Object.keys(data).map(
        key => `${encodeURIComponent(key)}=${encodeURIComponent(data[key])}`
      ).join('&');
    },
    submitForm() {
      axios.post(
        '/',
        this.encode({
          'form-name': 'contact',
          ...this.form
        }),
        {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
          }
        }
      ).then(() => {
        this.createNotification(0, 'Thanks for contacting me. I will get back to you as soon as I can!');
      }).catch(err => {
        this.createNotification(1, `Could not send message: ${err}.`)
      }).then(() => {
        this.isHidden = true;

        this.form.name = '';
        this.form.email = '';
        this.form.message_body = '';
      });
    }
  }
}
</script>

<style lang="sass" scoped>
  .modal
    animation: window-fade-in 0.4s
    -moz-animation: window-fade-in 0.4s
    -webkit-animation: window-fade-in 0.4s
    -o-animation: window-fade-in 0.4s
</style>
