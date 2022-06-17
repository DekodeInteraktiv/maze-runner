import React, { useState, useEffect } from 'react';
import './App.css';
const api = 'https://maze.peterbooker.com/api/v1/game/';

function App() {
  const [id, setId] = useState("");
  const [pw, setPw] = useState("");
  const [name, setName] = useState("");
  const [auth, setAuth] = useState(false);
  useEffect(() => {
    if(auth) {
      setInterval(() => {
        info();
      }, 500);
    }
  },[auth])
  const register = () => {
    fetch(api + id + '/player/register/' + pw, {
      method: 'POST',
      body: JSON.stringify({name: name, color: '#000',styles:{
  head: 'background: red;',
  body: 'background: blue; width: 30%;',
  feet: 'width: 30%;',
  foot: 'background: orange;',
  arm:  'background: red;'
}})
    })
    .then(response => response.json())
    .then(data => {
      console.log('Registred game', data);
      setAuth(data.token);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  };
  const info = () => {
    var headers = new Headers();
    headers.append("Authorization", "Bearer " + auth);
    headers.append("Content-Type", "application/json");
    fetch(api + id + '/player/status', {
      method: 'GET',
      headers: headers
    })
    .then(response => response.json())
    .then(data => {
      console.log('Info', data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  };
  const move = (direction) => {
    var headers = new Headers();
    headers.append("Authorization", "Bearer " + auth);
    headers.append("Content-Type", "application/json");
    fetch(api + id + '/player/move', {
      method: 'POST',
      headers: headers,
      body: JSON.stringify({direction: direction, distance: 1})
    })
    .then(response => response.json())
    .then(data => {
      console.log('Move', data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  };
  const shoot = (direction) => {
    var headers = new Headers();
    headers.append("Authorization", "Bearer " + auth);
    headers.append("Content-Type", "application/json");
    fetch(api + id + '/player/ability/shoot', {
      method: 'POST',
      headers: headers,
      body: JSON.stringify({direction: direction})
    })
    .then(response => response.json())
    .then(data => {
      console.log('Shoot', data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  };
  const bomb = () => {
    var headers = new Headers();
    headers.append("Authorization", "Bearer " + auth);
    headers.append("Content-Type", "application/json");
    fetch(api + id + '/player/ability/bomb', {
      method: 'GET',
      headers: headers
    })
    .then(response => response.json())
    .then(data => {
      console.log('Bomb', data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  };
  const keyControllers = (e) => {
    if(e.key) {
      if(e.key === 'ArrowLeft') {
        move('west')
      }

      if(e.key === 'ArrowRight') {
        move('east')
      }

      if(e.key === 'ArrowUp') {
        move('north')
      }

      if(e.key === 'ArrowDown') {
        move('south')
      }
    }
  }
  return (
    <div className="App">
    <div>
      <input value={id} onChange={e => setId(e.target.value)} placeholder="game id" />
      <input value={pw} onChange={e => setPw(e.target.value)} placeholder="game password" />
      <input value={name} onChange={e => setName(e.target.value)} placeholder="player name" />
      <button onClick={register}>Register</button>
    </div>
    {auth &&
      <div>
      <h2>Registred</h2>
      <div><button onClick={() => move('north')}>▲</button></div>
      <button onClick={() => move('west')}>⮜</button>
      <button onClick={() => move('south')}>▼</button>
      <button onClick={() => move('east')}>➤</button>
      <div><input type="text" value="" onKeyDown={keyControllers} placeholder="Select to navigate with keyboard" /></div>
      <div><button onClick={bomb}>Bomb</button></div>
      <h3>Shoot</h3>
        <div><button onClick={() => shoot('north')}>▲</button></div>
        <button onClick={() => shoot('west')}>⮜</button>
        <button onClick={() => shoot('south')}>▼</button>
        <button onClick={() => shoot('east')}>➤</button>
      </div>
    }
    </div>
  );
}

export default App;
