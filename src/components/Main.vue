<template>
  <div>
    <div v-for="i in numRows" v-bind:key="i" class="tile is-ancestor">
      <div v-for="card in cards.slice((i - 1) * 3, i * 3)" v-bind:key="card" class="tile is-4 is-parent">
        <component v-bind:is="card"></component>
      </div>
    </div>
  </div>
</template>


<script>
import AboutMe from './cards/AboutMe.vue'
import TitleCard from './cards/TitleCard.vue'

export default {
  name: 'Main',
  components: {
    TitleCard,
    AboutMe
  },
  computed: {
    cards: function () {
      // Return a list of ordered components to render.
      const keys = Object.keys(this.$options.components)
      return keys.filter(e => e != this.$options.name)
    },
    numRows: function () {
      return Math.ceil(this.cards.length)
    }
  }
}
</script>
