import useSWR from 'swr'
import './App.css'
// import { Box } from '@mantine/core'


export const ENDPOINT = "http://localhost:3000";

const fetcher = (url: string) => fetch(`${ENDPOINT}/${url}`).then(res => res.json())


function App() {
  const { data, mutate } = useSWR('api/todos', fetcher)
  return <h1>{JSON.stringify(data)}</h1>
}

export default App
