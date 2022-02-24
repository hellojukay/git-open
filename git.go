package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func GitRemotes() ([]string, error) {
	bin, err := exec.LookPath("git")
	if err != nil {
		return nil, fmt.Errorf("git command not found ? , %s", err)
	}
	c := exec.Command(bin, "remote", "-v")
	b, err := c.Output()
	if err != nil {
		return nil, fmt.Errorf("can not read git remote url , %s", err)
	}
	if string(b) == "" {
		return nil, fmt.Errorf(`remote not found, please run "git remote add" first`)
	}
	lines := strings.Split(string(b), "\n")
	var result []string
	re := regexp.MustCompile(`[\t\s]+`)
	for _, line := range lines {
		if line == "" {
			continue
		}

		arr := re.Split(line, -1)
		if len(arr) == 3 {
			result = append(result, arr[1])
		}
	}
	return git2https(uniq(result)), nil
}

//remove duplicate string
func uniq(list []string) []string {
	var m = make(map[string]bool)
	for i := range list {
		m[list[i]] = true
	}
	var result []string
	for k := range m {
		result = append(result, k)
	}
	return result
}

//git ssh protocol address convert to https protocol address
//$remote = $remote =~ s/\.git$//r;
//$remote = $remote =~ s/^git@/https:\/\//r;
//$remote = $remote =~ s/(:)([^\/])/\/$2/r;
func git2https(origins []string) []string {
	var result []string
	end := regexp.MustCompile(`.git$`)
	protocol := regexp.MustCompile(`^git@`)
	s := regexp.MustCompile(`(:)([^\/])`)
	for _, origin := range origins {
		r := end.ReplaceAll([]byte(origin), []byte(""))
		r = protocol.ReplaceAll(r, []byte("https://"))
		r = s.ReplaceAll(r, []byte("/$2"))
		result = append(result, string(r))
	}
	return result
}

func ISGithub(origin string) bool {
	return strings.Contains(origin, "https://github.com")
}

func WithPipeline(origin string) string {
	if ISGithub((origin)) {
		return origin + `/actions`
	}
	return origin + `/-/pipelines`

}
