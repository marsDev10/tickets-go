
// Components
import Header from "./components/Header"
import Modals from "./components/Modals"
import TeamList from "./components/TeamList"

const Teams = () => {
  return (
    <section className="flex flex-col gap-8 p-8 text-white">
        <Header/>
        <TeamList/>
        <Modals/>
    </section>
  )
}

export default Teams