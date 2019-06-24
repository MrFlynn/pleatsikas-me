<template>
  <div class="tile is-ancestor">
    <div class="tile is-4 is-parent">
      <div class="tile is-child box">
        <figure class="image is-square">
          <img src="../assets/profile.jpg" alt="Profile image.">
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
    </div>
    <div class="tile is-4 is-parent">
      <div class="tile is-child box">
        <h1 class="title is-3">About Me</h1>
        <p>
          Hello there, and welcome to my website! I'm Nick:
          an undergraduate student at the 
          <a href="https://ucr.edu">University of California, Riverside</a> 
          studying Computer Engineering. I've spent the last three years
          working at startup focused on deep learning running their server 
          infrastructure as well as writing developer tools and business
          automation software. Read more about my experience in my resume
          linked to the left.
        </p>
      </div>
    </div>
  </div>
</template>


<script>
export default {
    name: 'Home'
}
</script>
