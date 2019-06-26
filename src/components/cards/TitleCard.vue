<template>
  <div class="tile is-child box">
    <figure class="image is-square">
      <img src="../../assets/profile.jpg" alt="Profile image.">
    </figure>
    <br>
    <div class="media-content">
      <p class="title is-3">Nick Pleatsikas</p>
      <p class="subtitle is-6">Software Engineer</p>
    </div>
    <br>
    <nav class="breadcrumb has-bullet-separator" aria-label="breadcrumbs">
      <ul>
        <li><a href="https://github.com/MrFlynn">Github</a></li>
        <li><a href="https://linkedin.com/in/pleatsikas">LinkedIn</a></li>
        <li><a href="resume.pdf">Resume</a></li>
      </ul>
    </nav>
  </div>
</template>

<script>
export default {
  name: 'TitleCard'
}
</script>
