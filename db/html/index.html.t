<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">

    <!-- css -->
    <link rel="stylesheet" href="css/pure-min.css">
    <link rel="stylesheet" href="css/app.css">

    <!-- javascript -->
    <script src="js/config.js"></script>
    <script src="js/app.js"></script>
<! --<script type="module">
  import { app } from './js/config.js'; // Or it could be simply `hello.js`
  hello('world');
</script> -->

    <title>World Database</title>
</head>
<body>
    <div class="pure-g">
        <div class="pure-u-1-8"></div>
        <div class="pure-u-3-4">
            <h1>World Database</h1>
            <form class="pure-form" onsubmit="event.preventDefault(); app.searchClick()">
                <fieldset>
                    <input id="city-input" type="text" placeholder="City" required>
                    <button type="submit" class="pure-button pure-button-primary">Search</button>
                </fieldset>
            </form>

            <table class="pure-table pure-table-bordered pure-table-striped hidden" id="results-table">
                <thead>
                    <th>Name</th>
                    <th>District</th>
                    <th>Country</th>
                    <th>Population</th>
                </thead>
                <tbody id="results-table-body">
                </tbody>
            </table>
        </div>
        <div class="pure-u-1-8"></div>
    </div>
</body>
</html>
