const { defineConfig } = require('cypress')

module.exports = defineConfig({
  cucumber: {
    features: './features',
    stepDefinitions: './stepDefinitions',
  },
})
