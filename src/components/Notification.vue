<template>
  <div v-if="!notificationHidden" :class="this.color">
    <button @click="notificationHidden = true" class="delete"></button>
    {{ this.notificationData.message }}
  </div>
</template>

<script>
import {eventBus} from '../main';

export default {
  name: 'Notification',
  data() {
    return {
      notificationData: {
        type: 0,
        message: ''
      }
    }
  },
  props: {
    notificationHidden: {
      type: Boolean,
      default: true
    }
  },
  computed: {
    title: function() {
      return this.notificationData.type == 0 ? 'Success' : 'Error';
    },
    color: function() {
      return 'notification ' + (this.notificationData.type == 0 ? 'is-success' : 'is-danger');
    }
  },
  created() {
    eventBus.$on('notificationHidden', (message) => {
      this.notificationData.type = message.type;
      this.notificationData.message = message.message;

      this.notificationHidden = false;
    });
  }
}
</script>

<style lang="sass">
  .notification
    position: fixed !important
    width: 25%
    bottom: 0
    left: 0
    margin-bottom: 2rem
    margin-left: 2rem
    z-index: 1
    animation: window-fade-in 0.4s
    -moz-animation: window-fade-in 0.4s
    -webkit-animation: window-fade-in 0.4s
    -o-animation: window-fade-in 0.4s
</style>