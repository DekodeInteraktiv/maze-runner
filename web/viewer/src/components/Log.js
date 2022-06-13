import SoundBoard from './SoundBoard';

const Prompt = ({logPrompt}) => {
  switch (logPrompt.type) {
    case 1:
      return (<div class="log"><span>Hurry up!</span></div>)
      break;
    case 2:
      return (<div class="log"><span>Game over.</span></div>)
      break;
    case 6:
      return (<div class="log"><span>Fatality!</span></div>)
      break;
    default:

  }
  return null;
}

const Log = ({log}) => {
  return (
    <>
      {log.map(logPrompt =>
        <Prompt logPrompt={logPrompt} />
      )}
    </>
  );
}

export default Log;
