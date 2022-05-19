import React, { useState, useEffect } from 'react';
import {
  useParams,
  useNavigate,
} from "react-router-dom";
import Maze from './components/Maze';
import Object from './components/Object';
import PlayerArea from './components/PlayerArea';
import Timer from './components/Timer';
import logo from './logo.png';
import './App.css';
const api = 'https://maze.peterbooker.com/api/v1/game/';
const roundTime = 120;
const size = 16;

function App() {
  const [countDown, setCountDown] = useState(0);
  const [gameState, setGameState] = useState({});
  let updateGameStateInterval;
  let navigate = useNavigate();
  let params = useParams();
  const {game} = params;
  useEffect(() => {
    if(countDown > 0) {
      setTimeout(() => {
        setCountDown(countDown - 1);
      }, 1000);
    }
  },[countDown]);

  const startRound = () => {
    fetch(api + `${game}/start`, {
      method: 'GET',
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
  const startGame = () => {
    fetch(api + 'create', {
      method: 'POST',
      body: JSON.stringify({size: size, distribution: -0.3, timelimit: roundTime})
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
    fetch(api + `${game}/info`, {
      method: 'GET',
    })
    .then(response => response.json())
    .then(data => {
      //console.log(data);
      setGameState(data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  }
  useEffect(() => {
    if(game) {
      updateGameState();
      setInterval(updateGameState, 200);
    }
  },[game]);

  if(!game) {
    return (
      <div className="App select-screen">
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

  let {maze} = gameState;

  return (
    <div className="App">
      <style>{`body{--map-X: ${size};--map-Y: ${size};}`}</style>
      <div className="app-head">
        <PlayerArea id="3" claims={gameState.claims} players={gameState.players} password={gameState.password} />
        <div className="branding">
        <img src={logo} alt="" />
        <div style={{textAlign:'center'}}>Game {gameState.id}</div>
        </div>
        <PlayerArea id="4" claims={gameState.claims} players={gameState.players} password={gameState.password} />
      </div>
      <Maze maze={maze} claims={gameState.claims} countDown={countDown} players={gameState.players}/>
      <div className="app-foot">
        <PlayerArea id="1" claims={gameState.claims} players={gameState.players} password={gameState.password} />
        <div className="controls">
          {!gameState.active && <button onClick={startRound}>Start round</button>}
          <Timer timer={gameState.timer} active={gameState.active} roundTime={roundTime} />
        </div>
        <PlayerArea id="2" claims={gameState.claims} players={gameState.players} password={gameState.password} />
      </div>
    </div>
  );
}

export default App;
