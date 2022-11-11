import * as React from 'react'
import * as ReactDOM from 'react-dom'
import Calendar, { MonthView } from 'react-calendar'
import 'react-calendar/dist/Calendar.css'

function App() {
  const [meals, setMeals] = React.useState([])

  async function getMeals(from, to) {
    console.log('before')
    const payload = JSON.stringify({
      from: from.toJSON(),
      to: to.toJSON(),
    })
    console.log(payload)
    const res = await fetch (`http://localhost:8080/showMeals`, {
      method: 'POST',
      body: payload,
    })
    console.log('after')
    const body = await res.json()
    setMeals(body || [])
  }

  function onChange({action, activeStartDate, value, view}) {
    console.log('onActiveStartDateChange', action, activeStartDate, value, view)
    getMeals(activeStartDate, activeStartDate)
    console.log(getMeals)
  }

  return (
        <div>
          <h1>Meals list</h1>
          <Calendar onActiveStartDateChange={onChange} />
          <button onClick={() => addMeal()}>Add Meal</button>
          <MealsList meals={meals} />
        </div>
  )
}

function MealsList(props) {
  return <ul>
    {
      props.meals.map(function(meal) {
        return <li key={meal.id}>{JSON.stringify(meal)}</li>
      })
    }
  </ul>
}

ReactDOM.render(
  <App/>,
  document.getElementById('root')
)



//State / fonction / affichage