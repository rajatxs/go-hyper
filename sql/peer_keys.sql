CREATE TABLE IF NOT EXISTS `peer_keys`(
   addr CHAR(38),
   auth_key CHAR(87),
   exch_key CHAR(130),

   PRIMARY KEY(addr, auth_key, exch_key)
);

INSERT INTO `peer_keys` (addr, auth_key, exch_key) VALUES 
   (
      'ofWT4vZRP0YuW47BvYbt51SHirLb0A6U0WMvKg', 
      'BHXxCH8fJO_fTWlGR6EI0z3I6I4mMxbc16QNwAPAjLJnG57HxlWhvoFjovVoqdFHptaw7ize9gfSxhM0WLls_r8', 
      'BJ_mqQAOz5XLDWkNLlqOHpcj4vuplxUpOzdK0W10-ZQfw4x5CmLS1F19Bgug4V-vp1lrttclhRcSCbBMvD9j0fuBX4K0J1NgXgW6mINy0uvNeiBjEoofWSsHSo6vUPZ9Fw'
   );
