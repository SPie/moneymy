<template>
  <div>
    <bar-chart :chart-data="chartData" :options="options" />
  </div>
</template>

<script>
import BarChart from '@/components/BarChart'
import {expensesPerYear} from '@/api/expenses'

export default {
  name: 'BarCharts',
  components: {BarChart},
  comments: {
    BarChart
  },
  data () {
    return {
      chartData: {},
      options: {
        responsive: true,
        scales: {
          xAxes: [{
            stacked: true,
          }],
          yAxes: [{
            stacked: true,
          }]
        }
      },
      colorMap: new Map(),
      colors: [
          '#bada55',
          '#7fe5f0',
          '#ff0000',
          '#ff80ed',
          '#407294',
          '#ffd700',
          '#e6e6fa',
          '#00ffff',
          '#ffa500',
          '#f7347a',
          '#ffff00',
          '#00ff00',
          '#ffc3a0',
          '#4ca3dd',
          '#ff7f50',
          '#468499',
          '#008000',
          '#660066',
          '#daa520',
          '#000080',
          '#dddddd',
          '#8b0000',
          '#f5f5f5',
          '#7fffd4',
          '#20b2aa',
          '#ff0000',
          '#00aa00',
          '#0000ff',
          '#990000',
          '#004400',
          '#55ff99',
          '#aaaa00'
      ]
    }
  },
  mounted () {
    this.getExpensesPerYears()
  },
  methods: {
    getExpensesPerYears () {
      expensesPerYear().then((response) => {
        this.colorMap = new Map()
        let i = 0

        let labels = []
        let categoryMap = new Map()
        response.data.expenses.forEach(expense => {
          if (!labels.includes(expense.date)) {
            labels.push(expense.date)
          }
          if (categoryMap.get(expense.category) === undefined) {
            categoryMap.set(expense.category, [])
          }
          if (this.colorMap.get(expense.category) === undefined) {
            this.colorMap.set(expense.category, this.colors[i])
            i = (i + 1) % this.colors.length
          }

          categoryMap.get(expense.category).push(expense.amount)
        })

        let datasets = []
        categoryMap.forEach((value, key) => {
          datasets.push({
            label: key,
            data: value,
            backgroundColor: this.colorMap.get(key)
          })
        })
        this.chartData = {
          labels: labels,
          datasets: datasets
          // datasets: response.data.expenses.map(expense => {
          //   if (this.colorMap.get(expense.category) === undefined) {
          //     this.colorMap.set(expense.category, this.colors[i])
          //     i = (i + 1) % this.colors.length
          //   }
          //
          //   return {
          //     label: expense.category,
          //     data: [expense.amount],
          //     backgroundColor: this.colorMap.get(expense.category),
          //     stack: expense.date
          //   }
          // })
        }
      })
    }
  }
}
</script>