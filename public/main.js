var app = new Vue({
  el: '#app',
  data: {
    stepper: 0,
    questionnaire: [],
    selectedAnswers: {},
    result: {},
    isLoading: true,
    showQuestionnaire: true,
    showResult: false,
  },
  beforeMount: function () {
    this.loadQuestionnaire()
  },
  methods: {
    submitAnswers: function (msg) {
      this.isLoading = true
      this.showQuestionnaire = false
      console.log(msg)
      console.log(this.selectedAnswers)
      axios.post('/questionnaire', {selectedAnswers: this.selectedAnswers})
        .then(res => {
          this.showResult = true
          this.isLoading = false
          this.result = res.data.questionnaireResult ? res.data.questionnaireResult : {}
        })
        .catch(e => this.failed('Unsuccesful post'))
    },
    loadQuestionnaire: function () {
      this.isLoading = true
      this.showResult = false
      this.showQuestionnaire = false
      this.stepper = 0

      axios.get('/questionnaire')
        .then(res => {
          this.questionnaire = res.data.items ? res.data.items : []
          this.isLoading = false
          this.showQuestionnaire = true
          for (q of this.questionnaire) {
            this.selectedAnswers[q.id] = undefined
          }
        })
        .catch(e => this.failed('Unsuccesful get'))
    }
  },
})
