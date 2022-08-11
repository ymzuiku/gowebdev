v=v0.1.3
git tag $v
git push --tags
go install github.com/ymzuiku/webdev@$v
echo "done."