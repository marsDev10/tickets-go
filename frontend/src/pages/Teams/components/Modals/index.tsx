import { useTeamContext } from '../../context/TeamsProvider';
import CreateTeam from '../Forms/CreateTeam';
import GeneralModal from '../../../../components/Modals/GeneralModal';

const Modals = () => {

    const {
        state: { showCreateTeam },
        setters: { setShowCreateTeam },
    } = useTeamContext();

  return (
        <GeneralModal
            isOpen={showCreateTeam}
            onClose={() => setShowCreateTeam(false)}
            title="Create New Team"
        >
            <CreateTeam/>
        </GeneralModal>
  )
}

export default Modals