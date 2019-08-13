<template>
  <div class="container">
    <h1 class="title is-3">Work experience</h1>
    <div class="content">
      <div v-for="(job, idx) in jobs" :key="idx">
        <div class="job-entry">
          <p>
            <span class="tag is-medium is-primary">{{ job.start }} - {{ job.end }}</span>
            <span class="tag is-medium is-dark">{{ job.position }}</span>
            <span class="tag is-medium is-info link">
              <a :href=job.company.url>{{ job.company.name }}</a>
            </span> 
          </p>
          <p><strong>Description:</strong> {{ job.description }}</p>
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
  @import "bulma/sass/utilities/initial-variables.sass"

  .job-entry
    padding-bottom: 1rem

    span
      margin-right: 0.4rem
      @media screen and (max-width: $tablet)
        margin-bottom: 0.4rem

      a
        color: white

    p
      margin-bottom: 0 !important

  .link
    &:hover
      background-color: $blue
</style>
