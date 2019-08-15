<template>
  <div class="container">
    <h1 class="title is-3">Work experience</h1>
    <div class="content">
      <div v-for="(job, i) in jobs" :key="i">
        <div class="job-entry">
          <p>
            <span class="tag is-medium is-info link">
              <a :href=job.company.url>{{ job.company.name }}</a>
            </span> 
            <span class="tag is-medium is-dark">{{ job.position }}</span>
            <span class="tag is-medium is-light">{{ job.start }} - {{ job.end }}</span>
          </p>
          <strong>Description:</strong>
          <ul>
            <li v-for="(line, j) in job.description" :key="j">{{ line }}</li>
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

  .job-entry
    padding-bottom: 1.5rem

    span
      margin-right: 0.4rem
      @media screen and (max-width: $tablet)
        margin-bottom: 0.4rem

      a
        color: white

    p
      margin-bottom: 0 !important

    ul
      margin-top: 0

  .link
    &:hover
      background-color: $blue
</style>
