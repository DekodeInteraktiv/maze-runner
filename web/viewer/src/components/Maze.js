import Object from './Object';

function objectValues(obj) {
    let vals = [];
    for (const prop in obj) {
        vals.push(obj[prop]);
    }
    return vals;
}
const Color = ({color}) => {
  if(color) {
    return  <div className={`color color-${color}`} />;
  }
  return null;
}

const Cell = ({coords, players = [], cell, claims}) => {
  let object = {};
  if(1 === cell) {
    object = {type: "obstacle"};
  } else if(players) {
    players.forEach((player, i) => {
      if(coords.X === player.Pos.X && coords.Y === player.Pos.Y) {
        object = {type: "player", id:i};
      }
    });
  }

  return (
    <div className={`cell`} >
      <Color color={claims[coords.X][coords.Y]} />
      <Object object={object} />
    </div>
  );
}

const Maze = ({maze, countDown, players, claims}) => {
  return (
    <div className={`map ${countDown ? 'countdown' : ''}`}>
    {(0 !== countDown) &&
      <div className="countdown" key={'countdown-' + countDown}><span>{(countDown === 1) ? 'Go!' : countDown - 1}</span></div>
    }
    {objectValues(maze).map( (row, x ) =>
      <div className="row" key={'row-' + x}>
        {objectValues(row).map( (cell, y) =>
          <Cell coords={{X:x, Y:y}} cell={cell} claims={claims} players={players} key={'cell-' + y} />
        )}
      </div>
    )}
    </div>
  );
}

export default Maze;
