import React, { useState, useEffect, useRef } from 'react';
import {
  useParams,
  useNavigate,
} from "react-router-dom";
import Maze from './components/Maze';
import Object from './components/Object';
import PlayerArea from './components/PlayerArea';
import Log from './components/Log';
import Timer from './components/Timer';
import logo from './logo.png';
import './App.css';
import menuAudio from './audio/menu.wav';
import gameAudio from './audio/game.wav';
import SoundBoard from './components/SoundBoard';
const api = 'https://mazegame.dekodes.no/api/v1/game/';
const roundTime = 120;
const size = 16;

const Music = ({game, active, musicRef}) => {
  let music;
  if(!game) {
    return null;
  }
  if(active) {
    music = gameAudio;
  } else {
    music = menuAudio;
  }
  if(musicRef.current) {
    musicRef.current.volume = 0.2;
  }
  return (
    <audio ref={musicRef} src={music} autoPlay loop />
  );
}

function App() {
  const [countDown, setCountDown] = useState(0);
  const [gameState, setGameState] = useState({});
  const musicRef = useRef(null);

  let updateGameStateInterval;
  let navigate = useNavigate();
  let params = useParams();
  const {game} = params;
  useEffect(() => {
    if(countDown > 0) {
      if(1 === countDown) {
        musicRef.current.play();
      }
      setTimeout(() => {
        setCountDown(countDown - 1);
      }, 1000);
    }
  },[countDown]);

  const startRound = () => {
    musicRef.current.pause();
    var headers = new Headers();
    headers.append("Protected", "Protected");
    fetch(api + `${game}/start`, {
      method: 'GET',
      headers: headers,
    })
    .then(response => response.json())
    .then(data => {
      console.log('start', data);
      setCountDown(6);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  }
  const stopRound = () => {
    musicRef.current.pause();
    var headers = new Headers();
    headers.append("Protected", "Protected");
    fetch(api + `${game}/stop`, {
      method: 'GET',
      headers: headers,
    })
    .then(response => response.json())
    .then(data => {
      console.log('stop', data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  }
  const resetRound = () => {
    musicRef.current.pause();
    var headers = new Headers();
    headers.append("Protected", "Protected");
    fetch(api + `${game}/reset`, {
      method: 'GET',
      headers: headers,
    })
    .then(response => response.json())
    .then(data => {
      console.log('reset', data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  }
  const startGame = () => {
    fetch(api + 'create', {
      method: 'POST',
      body: JSON.stringify({
        size: size,
        distribution: -0.3,
        timelimit: roundTime,
        key:'SjqjcN81Shq77nqwLL',
        protected: true,
      })
    })
    .then(response => response.json())
    .then(data => {
      console.log('New game', data);
      navigate("/" + data.id);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  }
  const updateGameState = () => {
    var headers = new Headers();
    headers.append("Protected", "Protected");
    fetch(api + `${game}/info`, {
      method: 'GET',
      headers: headers,
    })
    .then(response => response.json())
    .then(data => {
      setGameState(data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  }
  useEffect(() => {
    if(game) {
      updateGameState();
      setInterval(updateGameState, 100);
    }
  },[game]);

  if(!game) {
    return (
      <div className="App select-screen">
        <Music game={game} active={gameState.active} musicRef={musicRef} />
        <div className="app-head">
          <div className="branding">
          <img src={logo} alt="" />
          </div>
        </div>
        <div className="intro-players">
        <Object object={{type: 'player', id:'default'}} />
        <Object object={{type: 'player', id:'default'}} />
        <Object object={{type: 'player', id:'default'}} />
        <Object object={{type: 'player', id:'default'}} />
        </div>
        <div className="app-foot">
          <div className="controls">
            <button onClick={startGame}>New game</button>
            <br />
            <button onClick={startGame}>Exit</button>
          </div>
        </div>
      </div>
    );
  }
  if(!gameState.maze) {
    return null;
  }

  const active = (gameState.status === 'running');

  let {maze} = gameState;
  return (
    <div className="App">
      <Music game={game} active={active} musicRef={musicRef} />
      <style>{`body{--map-X: ${maze.length};--map-Y: ${maze.length};}`}</style>
      <div className="app-head">
        <PlayerArea id="3" claims={gameState.claims} players={gameState.players} password={gameState.password} />
        <div className="branding">
        <img src={logo} alt="" />
        <div style={{textAlign:'center'}}>Game {gameState.id}</div>
        </div>
        <PlayerArea id="2" claims={gameState.claims} players={gameState.players} password={gameState.password} />
      </div>
      <Log log={gameState.log} />
      <Maze maze={maze} log={gameState.log} claims={gameState.claims} objects={gameState.objects} countDown={countDown} players={gameState.players}/>
      <div className="app-foot">
        <PlayerArea id="1" claims={gameState.claims} players={gameState.players} password={gameState.password} />
        <div className="controls">
          {!active && <button onClick={startRound}>Start round</button>}
          {active && <button onClick={stopRound}>Stop round</button>}
          {active && <button onClick={resetRound}>Reset round</button>}
          <Timer timer={gameState.timer} status={gameState.status} roundTime={roundTime} />
        </div>
        <PlayerArea id="4" claims={gameState.claims} players={gameState.players} password={gameState.password} />
      </div>
    </div>
  );
}

export default App;
