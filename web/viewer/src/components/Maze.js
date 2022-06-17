import Object from './Object';
import SoundBoard from './SoundBoard';

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

const Cell = ({coords, players = [], cell, claims, objects, log}) => {
  let object = {};
  let prop = {};
  let logEvent = {};
  if(1 === cell) {
    object = {type: "obstacle"};
  } else if(players) {
    players.forEach((player, i) => {
      if(JSON.stringify(coords) === JSON.stringify(player.pos)) {
        object = {type: "player", id:i};
      }
    });
  }
  if(objects) {
    objects.forEach((object) => {
      if(JSON.stringify(coords) === JSON.stringify(object.pos) ) {
        prop = {type: "prop", id:object.type};
      }
    });
  }
  if(log) {
    log.forEach((item) => {
      if(JSON.stringify(coords) === JSON.stringify(item.pos) ) {
        logEvent = {type: "event", id:item.type};
      }
    });
  }

  return (
    <div className={`cell`} >
      <Color color={claims[coords.X][coords.Y]} />
      <Object object={object} players={players} />
      <Object object={prop} />
      <Object object={logEvent} />
    </div>
  );
}

const Maze = ({maze, countDown, players, claims, objects, log}) => {
  return (
    <div className={`map ${countDown ? 'countdown' : ''}`}>
    {(0 !== countDown) &&
      <div className="countdown" key={'countdown-' + countDown}>
      <SoundBoard.Blip1 />
      <span>{(countDown === 1) ? 'Go!' : countDown - 1}</span>
      </div>
    }
    {objectValues(maze).map( (row, x ) =>
      <div className="row" key={'row-' + x}>
        {objectValues(row).map( (cell, y) =>
          <Cell coords={{X:x, Y:y}} cell={cell} objects={objects} log={log} claims={claims} players={players} key={'cell-' + y} />
        )}
      </div>
    )}
    </div>
  );
}

export default Maze;
