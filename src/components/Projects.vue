<template>
  <div class="container">
    <h1 class="title is-2">Projects</h1>
    <div class="project-entry content" v-for="(project, i) in projects" :key="i">
      <div class="field is-grouped is-grouped-multiline">
        <span class="tag is-large is-info link">
          <a :href=project.url v-if="project.url != ''">{{ project.name }}</a>
          <p v-else>{{ project.name }}</p>
        </span>
        <span class="tag is-large is-dark">Built with</span>
        <span v-for="(lang, j) in project.language" :key="j" class="tag is-large is-light">{{ lang }}</span>
      </div>
      <p class="project-description">{{ project.description }}</p>
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

  .field
    float: left
    margin: 0.1rem 0 0 0.5rem

  .project-description
    clear: both
    padding-left: 1rem

  .tag
    margin-bottom: 0.75rem
</style>