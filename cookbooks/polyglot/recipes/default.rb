ubuntu_mirror = 'http://archive.ubuntu.com/ubuntu/'
# ubuntu_mirror = 'http://mirror.rcg.sfu.ca/mirror/ubuntu/'
ubuntu_release = 'focal'
ubuntu_version = '20.04'
username = 'vagrant'
user_home = '/home/' + username
project_home = user_home + '/project/' # you may need to change the working directory to match your project


python3_packages = '/usr/local/lib/python3.8/dist-packages'
ruby_gems = '/var/lib/gems/2.7.0/gems/'


# Get Ubuntu sources set up and packages up to date.

template '/etc/apt/sources.list' do
  variables(
    :mirror => ubuntu_mirror,
    :release => ubuntu_release
  )
  notifies :run, 'execute[apt-get update]', :immediately
end
execute 'apt-get update' do
  action :nothing
end
execute 'apt-get upgrade' do
  command 'apt-get dist-upgrade -y'
  only_if 'apt list --upgradeable | grep -q upgradable'
end
directory '/opt'
directory '/opt/installers'


# Basic packages many of us probably want. Includes gcc C and C++ compilers.

#package ['build-essential', 'cmake']


# Other core language tools you might want

package ['python3', 'python3-pip', 'python3-dev', 'python3-virtualenv']  # Python
package ['golang-go']  # golang
execute 'pip3 install flask' do # Flask
  creates "#{python3_packages}/flask/__init__.py"
end
