import { useUsersContext } from '../../context/UsersProvider';

import GeneralModal from '../../../../components/Modals/GeneralModal';
import CreateUser from '../CreateUser';

const Modals = () => {
  
    const {
      state: {
        showCreateUser,
      },
      setters: {
        setShowCreateUser,
      }
    } = useUsersContext()
  
    return (
    <>
        <GeneralModal
          isOpen={showCreateUser}
          onClose={() => setShowCreateUser(false)}
          title="Create New User"
        >
          <CreateUser/>
        </GeneralModal>
    </>
  )
}

export default Modals;