# Automatisation des commits
while true; do
    git add .
    git commit -m "Basic Groupie Tracker"
    git push
    sleep 60
done