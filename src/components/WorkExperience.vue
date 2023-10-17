<template>
  <div class="container">
    <h1 class="title is-2">Work experience</h1>
    <div class="content">
      <div v-for="(job, i) in jobs" :key="i">
        <div class="job-entry">
          <div class="field is-grouped is-grouped-multiline">
            <span class="tag is-large is-info link">
              <a :href=job.company.url>{{ job.company.name }}</a>
            </span>
            <span class="tag is-large is-dark">{{ job.position }}</span>
            <span class="tag is-large is-light">{{ job.start }} ⇢ {{ job.end }}</span>
          </div>
          <ul>
            <li v-for="(line, j) in job.description" :key="j" v-html="line"></li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import experience from '../assets/experience.json'

export default {
  name: 'WorkExperience',
  computed: {
    jobs() {
      return experience;
    }
  }
}
</script>

<style lang="sass" scoped>
  $blue: #0E7C7B
  @import "bulma/sass/utilities/initial-variables.sass"

  .tag
    margin-top: 0.5rem

  .job-entry
    padding-bottom: 1.5rem

    p
      margin-bottom: 0 !important

    ul
      margin-top: 0.75rem
</style>
