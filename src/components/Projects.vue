<template>
  <div class="container">
    <h1 class="title is-3">Projects</h1>
    <div class="project-entry" v-for="(project, idx) in projects" :key="idx">
      <h4 class="project-title subtitle is-5">
        <a class="underline" :href=project.url v-if="project.url != ''">{{ project.name }}</a>
        <p v-else>{{ project.name }}</p>
      </h4>
      <div class="field is-grouped is-multiline">
        <div class="control">
          <div class="tags has-addons">
            <span class="tag is-dark">Language</span>
            <span class="tag is-info">{{ project.language }}</span>
          </div>
        </div>
      </div>
      <p class="project-description"><strong>Description: </strong>{{ project.description }}</p>
    </div>
  </div>
</template>

<script>
import info from '../assets/projects.json'

export default {
  name: 'Projects',
  computed: {
    projects() {
      return info;
    }
  }
}
</script>

<style lang="sass" scoped>
  .project-entry
    padding-bottom: 1.5rem

  .project-title
    margin-bottom: 0 !important
    float: left

  .field
    float: left
    margin: 0.1rem 0 0 0.5rem

  .project-description
    clear: both
    padding-left: 1rem
</style>