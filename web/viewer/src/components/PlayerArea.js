import SoundBoard from './SoundBoard';

const PlayerArea = ({id, players, password, claims}) => {
  if(!players || !players[id-1]) {
    return(
      <div key={`player-${id}`} className={`player-area player-area-${id}`}>
        <span className="name flicker">Insert coin(s)</span>
        {password && <span className="password">Password: <span>{password}</span></span>}
      </div>
    )
  }
  const player = players[id-1];
  let score = 0;
  claims.forEach((column) => {
    column.forEach((cell) => {
      if(parseInt(id) === cell) {
        score++;
      }
    });
  });

  return(
    <div key={`player-${id}`} className={`player-area player-area-${id}`}>
      <SoundBoard.Select />
      <span className="name">{player.name}</span>
      <span className="score">Score: {score}</span>
    </div>
  )
}

export default PlayerArea;
