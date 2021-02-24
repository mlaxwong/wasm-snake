const express = require('express')
const app = express()

app.engine('ejs', require('ejs-locals'))
app.set('view-engine', 'ejs')

app.get('/', (req, res) => {
  res.render('index.ejs', {user: {name: 'Mlax'}})
})

app.listen(8080)