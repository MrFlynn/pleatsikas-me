<template>
  <div id="app">
    <div class="columns is-gapless">
      <div class ="column is-one-quarter has-background-light">
        <h1>Hello world</h1>
      </div>
      <div class="column">
        <h1>This is the content region</h1>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style lang="sass">
  @import "bulma"
</style>

<style>
  .columns {
    display: flex;
    height: 100vh;
  }
</style>
