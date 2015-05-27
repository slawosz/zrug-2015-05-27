require 'test_helper'

class WorkersControllerTest < ActionController::TestCase
  test "should get trigger" do
    get :trigger
    assert_response :success
  end

end
