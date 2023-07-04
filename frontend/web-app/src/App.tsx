import React from 'react';
import './src/css/main.css';

function App() {
  return (
	  <div id="wrapper">
			<div id="bg"></div>
			<div id="overlay"></div>
			<div id="main">

				<!-- Header -->
					<header id="header">
						<h1>DevOps Driver</h1>
						<p>If you have or need the words below <span style="font-weight:bold;font-style:italic;">you're a fit.</span>
						<br>Click on them to see what we can do or simply <span style="font-weight:bold">contact us</span>
						<br>for a completely free and extensive consultation.</span></p>
						<nav>
						    <ul>
						        <li><a href="#">aws</a></li>
						        <li><a href="#">terraform</a></li>
						        <li><a href="#">cicd</a></li>
						        <li><a href="#">kubernetes</a></li>
						        <li><a href="#">functions</a></li>
						        <li><a href="#">serverless</a></li>
						    </ul>
						    <ul>
						        <li><a href="#">iac</a></li>
						        <li><a href="#">cloud</a></li>
						        <li><a href="#">observability</a></li>
						        <li><a href="#">cost</a></li>
						        <li><a href="#">security</a></li>
						        <li><a href="#">database</a></li>
						    </ul>
						    <ul>
						        <li><a href="#">performance</a></li>
						        <li><a href="#">loadtesting</a></li>
						        <li><a href="#">chaos</a></li>
						        <li><a href="#">lift</a></li>
						    </ul>
						</nav>
					</header>

				<!-- Footer -->
					<footer id="footer">
						<!--<span class="copyright">&copy; Untitled. Design: <a href="http://html5up.net">HTML5 UP</a>.</span>-->
					</footer>

			</div>
		</div>
  );
}

export default App;
