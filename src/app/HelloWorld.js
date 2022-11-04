import * as React from 'react'
import * as ReactDOM from 'react-dom'

function App() {
  const [meals, setMeals] = React.useState([])

  async function getMeals() {
    const res = await fetch (`http://localhost:8080/showMeals`, {
      method: 'POST',
      body: JSON.stringify({
        from: "2020-01-01T00:00:00Z",
        to: "2024-01-01T00:00:00Z"
      })
    })
    const body = await res.json()
    setMeals(body)
  }

  return (
        <div>
          <h1>Meals list</h1>
          <button onClick= {() => getMeals()}>Get meals</button>
          <MealsList meals={meals} />
        </div>
  )
}

function MealsList(props) {
  return <ul>
    {
      props.meals.map(function(meal) {
        return <li>{JSON.stringify(meal)}</li>
      })
    }
  </ul>
}

ReactDOM.render(
  <App/>,
  document.getElementById('root')
)



//State / fonction / affichage