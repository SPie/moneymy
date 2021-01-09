import client from './client'

const expensesPerYear = () => client.get('/years')

export {expensesPerYear}