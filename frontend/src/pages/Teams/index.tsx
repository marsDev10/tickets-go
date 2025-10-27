
// Components
import Header from "./components/Header"
import TeamList from "./components/TeamList"

const Teams = () => {
  return (
    <section className="flex flex-col gap-8 p-8 text-white">
        <Header/>
        <TeamList/>
    </section>
  )
}

export default Teams