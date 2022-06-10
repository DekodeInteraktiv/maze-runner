import React, { useState, useEffect, useRef } from 'react';
import logo from './logo.png';
import './App.css';

function App() {
    const [showPassword, setShowPassword] = useState(false);
    const getTitle = () => {
      if (showPassword) {
        return "2:JsNh!7";
      } else {
        return 'LEADERBOARD';
      }
    }
    return (
      <div className="App select-screen">
        <div className="app-head">
          <div className="branding">
          <img onClick={() => setShowPassword(true)} src={logo} alt="" />
          <h1>{getTitle()}</h1>
          </div>
        </div>
        <div>
          <ul>
            <li class="head">
              <span>Player</span>
              <span>Wins</span>
            </li>
            <li>
              <span>$EliteHaxxor</span>
              <span>197</span>
            </li>
            <li>
              <span>$EliteHaxxor2</span>
              <span>133</span>
            </li>
            <li>
              <span>$EliteHaxxor3</span>
              <span>128</span>
            </li>
            <li>
              <span>%TotallyNotEliteHaxxor</span>
              <span>105</span>
            </li>
            <li>
              <span>DoomGuy</span>
              <span>27</span>
            </li>
            <li>
              <span>PartyParrot</span>
              <span>24</span>
            </li>
            <li>
              <span>PeterBooker</span>
              <span>23</span>
            </li>
            <li>
              <span>PizzaWill</span>
              <span>18</span>
            </li>
            <li>
              <span>Zpaceinvader</span>
              <span>18</span>
            </li>
            <li>
              <span>DrDisrespect</span>
              <span>17</span>
            </li>
            <li>
              <span>PewDiePie</span>
              <span>15</span>
            </li>
          </ul>
        </div>
      </div>
    );
}

export default App;
