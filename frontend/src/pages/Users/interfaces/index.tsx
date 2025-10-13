import { useUsersContext } from "../context/UsersProvider"

const Modal = () => {
    const {
        state: { showCreateUser },
        setters: { setShowCreateUser },
    } = useUsersContext();

    


    return (
    <div>Modal</div>
  )
}

export default Modal