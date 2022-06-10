import React, { useState, useEffect, useRef } from 'react';
import logo from './logo.png';
import './App.css';

function App() {
    const [screen, setScreen] = useState(1);
    const [fail, setFail] = useState(false);
    const [translate, setTranslate] = useState(0);
    const translateArray = [
      'translate(200%,200%)',
      'translate(200%,0)',
      'translate(0,200%)',
      'translate(0,0)',
      'translate(-200%,-200%)',
      'translate(-200%,0)',
      'translate(0,-200%)',
      'translate(0,0)'
    ]
    const [answer1, setAnswer1] = useState(0);
    const [answer2, setAnswer2] = useState(0);
    const [answer3, setAnswer3] = useState(0);
    const [emotion, setEmotion] = useState('');
    const getScreen = () => {
      switch (screen) {
        case 1:
          return(
            <div className="step step-1">
              <p>Please enter your email</p>
              <input type="text" placeholder="your@email.here" />
              <div><button onClick={() => setScreen(2)}>Submit</button></div>
            </div>
          );
          case 2:
            const translateCycle = () => {
              if(translate >= translateArray.length - 1) {
                setTranslate(0);
              } else {
                setTranslate(translate + 1);
              }

            }
            return(
              <div className="step step-2">
                <p>Lets just make sure you're not a robot</p>
                <div><button style={{transform: translateArray[translate]}} onMouseEnter={() => translateCycle()} onClick={() => setScreen(3)}>I'm not a robot</button></div>
              </div>
            );
          case 3:
            const verify = () => {
              if(answer1 === 1 && answer2 === 3 && answer3 === 1) {
                setScreen(4);
              } else {
                setFail(true);
              }
            }
            return(
              <div className="step step-3">
                <p>Wait how does clicking a button prove you're not a robot? Here try something harder</p>
                <div>
                  <h3>What is 6 + 5?</h3>
                  <div className="radio-buttons">
                    <label><input onClick={() => setAnswer1(1)} type="radio" name="question-1"></input>9</label>
                    <label><input onClick={() => setAnswer1(2)} type="radio" name="question-1"></input>12</label>
                    <label><input onClick={() => setAnswer1(3)} type="radio" name="question-1"></input>782</label>
                  </div>
                  <h3>What is 25 % 7?</h3>
                  <div className="radio-buttons">
                    <label><input onClick={() => setAnswer2(1)} type="radio" name="question-2"></input>38</label>
                    <label><input onClick={() => setAnswer2(2)} type="radio" name="question-2"></input>1</label>
                    <label><input onClick={() => setAnswer2(3)} type="radio" name="question-2"></input>4</label>
                  </div>
                  <h3>What is true + false?</h3>
                  <div className="radio-buttons">
                    <label><input onClick={() => setAnswer3(1)} type="radio" name="question-3"></input>true</label>
                    <label><input onClick={() => setAnswer3(2)} type="radio" name="question-3"></input>false</label>
                    <label><input onClick={() => setAnswer3(3)} type="radio" name="question-3"></input>NaN</label>
                  </div>
                  <div><button onClick={() => verify()}>Submit</button></div>
                </div>
              </div>
            );
          case 4:
          const verifyEmotion = () => {
            if(emotion !== '') {
              setScreen(5);
            } else {
              setFail(true);
            }
          }
           return(
             <div className="step step-4">
               <p>Wait a minute... that's exactly what a robot would answer isn't it?<br /> This will definitely fool you!</p>
               <div>
                 <h3>What is your favourite emotion</h3>
                 <input type="text" value={emotion} onChange={(e) => setEmotion(e.target.value)} placeholder="Insert human emotion here" />
                 <div><button onClick={() => verifyEmotion()}>😫 Submit 😍</button></div>
               </div>
             </div>
           )
          case 5:
          return(
            <div className="step step-5">
              <p>Congratulations, you are not a robot</p>
              <div>
                <h3>API key</h3>
                <input type="text" value="SjqjcN81Shq77nqwLL"/>
              </div>
            </div>
          )
      }
    }
    return (
      <div className="App select-screen">
        <div className="app-head">
          <div className="branding">
          <img src={logo} alt="" />
          </div>
        </div>
        <h2>Register API-key</h2>
          {(fail === true) &&
            <div><h2>YOU ARE NOT WORTHY!</h2></div>
          }
          {(fail === false) &&
            <div>{getScreen()}</div>
          }
      </div>
    );
}

export default App;
