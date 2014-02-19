var files = {
  "foundation_js": [
  "app/assets/js/foundation/foundation.abide.js",
  "app/assets/js/foundation/foundation.accordion.js",
  "app/assets/js/foundation/foundation.alert.js",
  "app/assets/js/foundation/foundation.clearing.js",
  "app/assets/js/foundation/foundation.dropdown.js",
  "app/assets/js/foundation/foundation.interchange.js",
  "app/assets/js/foundation/foundation.joyride.js",
  "app/assets/js/foundation/foundation.js",
  "app/assets/js/foundation/foundation.magellan.js",
  "app/assets/js/foundation/foundation.offcanvas.js",
  "app/assets/js/foundation/foundation.orbit.js",
  "app/assets/js/foundation/foundation.reveal.js",
  "app/assets/js/foundation/foundation.tab.js",
  "app/assets/js/foundation/foundation.tooltip.js",
  "app/assets/js/foundation/foundation.topbar.js"
  ],
  
  "kevin_sample_js": [
    "app/assets/js/kevin_sample/kevin_sample_util.js"
  ]
}

module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    
    jshint: {
      options: {
        jshintrc: 'app/assets/js/.jshintrc'
      },
      gruntfile: {
        src: 'Gruntfile.js'
      },
      src: {
        src: ['app/assets/js/**/*.js']
      }
    },

    uglify: {
      options: {
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n',
        sourceMap: true,
        report: 'min'
      },
      foundation_js: {
        files: {
          'public/js/foundation-custom.min.js': files.foundation_js
        }
      },
      kevin_sample_js: {
        files: {
          'public/js/kevin_sample.min.js': files.kevin_sample_js
        }
      }
    },

    sass: {
      foundation_css: {
        options: {
          style: 'compressed'
        },
        files: {
          'public/css/foundation-custom.min.css': 'app/assets/sass/foundation_custom.scss'
        }        
      }
      ,
      kevin_sample_css: {
        options: {
          style: 'expanded',
          lineNumbers: true
        },
        files: {
          'public/css/kevin_sample.css': 'app/assets/sass/kevin_sample.sass'
        }        
      }
    },


  });

  grunt.event.on('watch', function(action, filepath, target) {
    grunt.log.writeln(target + ': ' + filepath + ' has ' + action);
  });


  // Load the plugin that provides the tasks.
  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-sass');

  // Register tasks.
  grunt.registerTask('default', ['uglify','sass']);
  grunt.registerTask('test', ['jshint', 'qunit']);

};